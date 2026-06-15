package main

import (
	"html/template"
	"net/http"

	"ascii-art-web/perser"
)

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("input")
	banner := r.FormValue("banner")

	if _, err := perser.Validate(text); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fonts, err := perser.LoadBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	words := perser.SplitInput(text)
	result := perser.GenerateArt(fonts, words)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}