import React, { useState, useEffect } from "react";
import "./App.css";

function App() {
  const [smsList, setSmsList] = useState([]);

  useEffect(() => {
    fetchSMSList();
  }, []);

  const fetchSMSList = async () => {
    const response = await fetch("http://localhost:8080/getSMS");
    const data = await response.json();
    setSmsList(data);
  };

  return (
    <div className="App">
      <header>
        <h1>smstrap</h1>
      </header>
      <main>
        <h2>Incoming SMS Messages</h2>
        {smsList.length === 0 ? (
          <p>No messages available.</p>
        ) : (
          <ul>
            {smsList.map((sms, index) => (
              <li key={index}>
                <strong>{sms.phoneNumber}:</strong> {sms.body}
              </li>
            ))}
          </ul>
        )}
      </main>
    </div>
  );
}

export default App;
