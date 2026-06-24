package handlers

import (
	"ascii-web/parser"
	"fmt"
	"html/template"
	"net/http"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "not found", 404)
		return
	}

	asciitmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "error parsing template file", 500)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
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
