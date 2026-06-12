package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title   string
	Text    string
	Version string
	Docs    string
}

func root(w http.ResponseWriter, r *http.Request) {
	data := pageData{
		Title:   "Ascii Art Web",
		Text:    "Welcome to my super fine ASCII ART text generator website, just input text and get ASCII representation in seconds!",
		Version: "1.0.0",
		Docs:    "API documentation: chech docs page",
	}

	var temp, err = template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, data)
	log.Printf("request recived on %s", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)
	log.Println("vist http://127.0.0:8000")
	http.ListenAndServe(":8000", mux)
}
