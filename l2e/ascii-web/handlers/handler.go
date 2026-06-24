package handlers

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Input  string
	Banner string
	Result string
}

var tmpl = template.Must(template.ParseFiles("templates/home.html"))

func HomeHanler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "not found", 404)
		return
	}

	hometmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "error parsing template file", 500)
		return
	}
	hometmpl.Execute(w, nil)
}
