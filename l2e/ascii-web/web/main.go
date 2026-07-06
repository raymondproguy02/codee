package main

import (
	"ascii-web/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHanler)
	mux.HandleFunc("/health", handlers.HandleHealth)
	mux.HandleFunc("/about", handlers.HandleAbout)
	mux.HandleFunc("POST /ascii-art", handlers.AsciiHandler)
	mux.HandleFunc("/ascii-switch", handlers.SwitchHandler)

	log.Println("Server running on http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
