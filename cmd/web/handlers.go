package web

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/html/base.html", "templates/html/index.html",
		"templates/html/includes/headers.html", "templates/html/includes/footer.html",
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SelectOrderID(w http.ResponseWriter, r *http.Request) {

}
