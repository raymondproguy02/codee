package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title  string
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

func echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/echo" {
		http.Error(w, "Path not allowed", http.StatusBadRequest)
		return
	}

	text := r.FormValue("input")
	if text == "" {
		http.Error(w, "Empty input", http.StatusBadRequest)
		return
	}

	data := pageData{
		Input:  text,
		Result: text,
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not fonud", http.StatusNotFound)
		return
	}
	temp.Execute(w, data)
	log.Println("Request recieved at:", r.URL)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/echo", echo)

	log.Println("Server up and running, visit http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
}
