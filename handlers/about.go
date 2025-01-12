package handlers

import (
	"html/template"
	"net/http"
)

// AboutHandler handles the root route
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/about.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		// TODO handle error(s)
		return
	}
}
