package main

import (
	"ascii-web/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHanler)

	log.Println("Server running on http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
