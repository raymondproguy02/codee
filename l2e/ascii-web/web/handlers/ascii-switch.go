package handlers

import (
	"ascii-web/parser"
	"fmt"
	"html/template"
	"net/http"
)

func SwitchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-switch" {
		http.Error(w, "not found", 404)
		return
	}

	asciitmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "error parsing template file", 500)
		return
	}
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	font, err := parser.LoadBanner(banner)
	if err != nil {
		http.Error(w, "error loading banner", 500)
		return
	}
	valid, err := parser.Val(text)
	if err != nil {
		fmt.Fprintf(w, "invalid character: %c", valid)
		http.Error(w, "invalid character", 400)
		return
	}
	words := parser.Split(text)
	res := parser.GenArt(font, words)

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: res,
	}
	asciitmpl.Execute(w, data)
}
