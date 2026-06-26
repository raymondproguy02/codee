package main

import (
	"html/template"
	"log"
	"net/http"
	"reloaded-web/processors"
)

type pageData struct {
	Input  string
	Result string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Path not allowed", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not fonud", http.StatusNotFound)
		return
	}
	temp.Execute(w, nil)
	log.Println("Request recieved at:", r.URL)
}

func style(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, "templates/style.css")
}

func reload(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("input") 
	res := processors.Hex(text)
	res = processors.Bin(text)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := pageData{
		Input:  text,
		Result: res,
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/style.css", style)
	mux.HandleFunc("/reload", reload)

	log.Println("Server up and running, visit http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
}
