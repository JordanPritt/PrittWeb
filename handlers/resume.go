package handlers

import (
	"html/template"
	"net/http"
)

// ResumeHandler handles the root route
func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/resume.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		// TODO handle error(s)
		return
	}
}
