package main

import (
	"html/template"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Path not allowed", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "Frontend page missing", http.StatusNotFound)
		return
	}
	temp.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
}
