package main

import (
	"log"
	"net/http"
	"text/template"
)

func root(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Println(err)
	}
	temp.Execute(w, nil)
}

func cssFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "text/css")
	http.ServeFile(w, r, "public/style.css")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	mux.HandleFunc("/style.css", cssFile)
}
