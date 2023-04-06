#!/bin/bash

# Change to the frontend directory
cd frontend

# Install npm dependencies
echo "Installing npm dependencies..."
npm install

# Build the React frontend
echo "Building the React frontend..."
npm run build

# Change back to the project root directory
cd ..

# Create the static folder if it doesn't exist
mkdir -p static

# Remove the existing frontend folder if it exists
rm -rf static/frontend

# Copy the build output to the static folder
echo "Copying the build output to the static folder..."
cp -r frontend/build static/frontend

# Build the Go executable
echo "Building the Go executable..."
go build -o smstrap

echo "Build and copy complete!"
