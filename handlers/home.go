package handlers

import (
	"html/template"
	"net/http"
)

// HomeHandler handles the root route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		// TODO handle error(s)
		return
	}
}
