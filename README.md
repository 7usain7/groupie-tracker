## Groupie Tracker:

roupie Tracker is a web application built with Go (Golang) that uses the (Groupie-Trackers API) to display information about
music artists, including their members, creation dates, first album, concert locations, and more.

## Features:

- Browse a full list of artists and bands  
- View artist details (name, members,first album, creation date, etc.)  
- Query to search artists by name, member, creation year, first album, or location
- Error handling and custom error pages


## Technologies Used:

- **Go (Golang)** — Backend logic, API calls, and HTTP server  
- **HTML/CSS** — Frontend rendering and styling  
- **Docker** — For containerizing the app  

## How to Run:

* Using Go: 

1. `go run main.go`

2. open the browser at: http://localhost:8080

* Using Docker:

1. `docker build -t groupie-tracker .`

2. `docker run -p 8080:8080 groupie-tracker`  

3.  open the browser at: http://localhost:8080


## Team Members:

- Ali Hussain #alimadan
- Hussain Abdulrasool #habdulras
