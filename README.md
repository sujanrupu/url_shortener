# URL Shortener
This is a simple URL shortener built using Go for the backend and HTML, JavaScript, and Tailwind CSS for the frontend.

## Features:
- Shortens long URLs into compact, shareable links.
- Provides a user-friendly interface with Tailwind CSS.
- Displays the shortened URL with a copy-to-clipboard button.
- Shows a loading animation while processing the request.
- Backend supports CORS for smooth frontend integration.

## Technologies :
### Backend:
- Go (Golang)
- net/http for handling HTTP requests
- encoding/json for JSON parsing
- url package for encoding URLs
### Frontend:
- HTML, JavaScript, and Tailwind CSS
- Fetch API for making requests to the backend
- Copy-to-clipboard functionality using JavaScript

## Installation:
Ensure Go is installes in the system. 
### Steps to Run the Project:
1. Clone the repository:
```
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```
2. Run the backend server:
```
go run main.go
```
3. Open index.html in browser or use a local web server to serve the frontend.
