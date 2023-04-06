# SMS Simulator

SMS Simulator is a simple web application that simulates the sending and receiving of SMS messages. The backend is built with Go, while the frontend uses React.

## Features

- Exposes an API to send SMS messages
- Stores SMS messages in a file (no database required)
- Displays incoming SMS messages in a React-based web interface

## Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or later
- [Node.js](https://nodejs.org/en/download/) and npm (included with Node.js)

## Installation

1. Clone the repository:

```
git clone https://github.com/aliharis/smstrap.git
cd smstrap
```

2. Install frontend dependencies and build the frontend:

```
cd frontend
npm install
npm run build
cd ..
```

3. Build the Go executable:

`go build -o sms-simulator`

## Running the Application

1. Run the Go server:

`./smstrap`

2. Access the application at `http://localhost:8080`.

## Testing

To test the SMS sending functionality, use a tool like `curl` or Postman to send a POST request to the `/sendSMS` endpoint:

`curl -X POST -H "Content-Type: application/json" -d '{"phoneNumber": "1234567890", "body": "Hello, World!"}' http://localhost:8080/sendSMS`

After sending an SMS, refresh the web interface, and you should see the new message appear in the list.

## Building the Bundled Executable

To build a bundled executable containing both the frontend and backend, run the provided build script:

This script will:

1. Install the npm dependencies for the frontend
2. Build the React frontend
3. Copy the frontend build output to the `static` folder in the project root directory
4. Build the Go executable, embedding the frontend files

After running the script, you'll have an updated `sms-simulator` executable that includes the latest frontend build.

To run the bundled executable, simply

`./smstrap`

## License

This project is released under the [MIT License](LICENSE).
