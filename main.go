package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

//go:embed static/frontend/*
var frontend embed.FS

type SMS struct {
	PhoneNumber string `json:"phoneNumber"`
	Body        string `json:"body"`
}

var smsList []SMS
var smsMutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/sendSMS", sendSMSHandler)
	http.HandleFunc("/getSMS", getSMSHandler)
	http.HandleFunc("/", frontendHandler)

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sendSMSHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var sms SMS
	if err := json.NewDecoder(r.Body).Decode(&sms); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	smsMutex.Lock()
	smsList = append(smsList, sms)
	saveSMSListToFile()
	smsMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func getSMSHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	smsMutex.Lock()
	defer smsMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(smsList)
}

func saveSMSListToFile() {
	data, err := json.Marshal(smsList)
	if err != nil {
		log.Println("Error marshalling SMS list:", err)
		return
	}

	if err := ioutil.WriteFile("sms_data.json", data, 0644); err != nil {
		log.Println("Error writing SMS data to file:", err)
	}
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("static/frontend", r.URL.Path)

	if filepath.Ext(path) == "" {
		path = filepath.Join(path, "index.html")
	}

	data, err := frontend.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(path))
	w.Header().Set("Content-Type", contentType)
	w.Write(data)
}

func init() {
	if _, err := os.Stat("sms_data.json"); err == nil {
		data, err := ioutil.ReadFile("sms_data.json")
		if err != nil {
			log.Println("Error reading SMS data from file:", err)
		} else {
			json.Unmarshal(data, &smsList)
		}
	}
}
