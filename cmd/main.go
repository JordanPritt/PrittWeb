package main

import (
	"fmt"
	"github.com/JordanPritt/PrittWeb/handlers"
	"net/http"
)

func main() {
	// Serve static files (like HTML, CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle the root route
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/resume", handlers.ResumeHandler)

	// Handle the HTMX request
	http.HandleFunc("/update", handlers.UpdateHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// TODO handle the error(s)
		return
	}
}
