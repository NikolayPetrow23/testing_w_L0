package router

import (
	"html/template"
	"log"
	"net/http"
	"testing_w_L0/my_cache"
)

func IndexRouter(w http.ResponseWriter, r *http.Request) {
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

	orderUid := r.FormValue("orderUid")

	context := map[string]interface{}{
		"OrderNotFound": orderUid,
	}

	if err := tmpl.Execute(w, context); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func OrderRouter(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/html/base.html", "templates/html/order.html",
		"templates/html/includes/headers.html", "templates/html/includes/footer.html",
	}

	tmpl, err := template.ParseFiles(files...)

	orderUid := r.FormValue("orderUid")

	data := my_cache.GetOrderFromCache(orderUid)

	if data == nil {
		http.Redirect(w, r, "/?error=OrderNotFound&orderUid="+orderUid, http.StatusSeeOther)
		return
	}

	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
