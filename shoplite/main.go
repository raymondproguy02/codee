package main

import (
	"html/template"
	"log"
	"net/http"
	"shoplite/handlers"
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

func styleFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "public/styles/style.css")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/styles/style.css", styleFile)
	mux.HandleFunc("/listings", handlers.GetHandler)
	mux.HandleFunc("api/add", handlers.PostHandler)
	log.Println("Server running on http:127.0.0.1:8000")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Panic(err)
	}
}
