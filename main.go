package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	PageTitle      string
	Username       string
	WelcomeMessage string
	Balance        string
	GameURL        string
}

func root(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	data := pageData{
		PageTitle:      "Golang billionaire",
		Username:       "Anonymous",
		WelcomeMessage: "Welcome to the official Golang Billionaire Game!",
		Balance:        "$70",
		GameURL:        "https://www.gobillionaire.dev",
	}

	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		fmt.Printf("Template parsing error: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = temp.Execute(w, data)
	if err != nil {
		fmt.Printf("Template execution error: %v\n", err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)

	fmt.Println("Server running on PORT 8000. Open http://localhost:8000 in your browser.")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
