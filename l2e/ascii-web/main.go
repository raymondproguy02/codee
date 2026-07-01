package main

import (
	"ascii-web/handlers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "OK",
	})
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := map[string]string{
		"project":     "01Edu Project",
		"name":        "Ascii Art Web",
		"version":     "1.0.0",
		"description": "Web version of ascii art!",
		"author":      "raymondproguy",
	}

	fJson, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintln(w, string(fJson))
	//json.NewEncoder(w).Encode(string(fJson))
	
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHanler)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("POST /ascii-art", handlers.AsciiHandler)
	mux.HandleFunc("/ascii-switch", handlers.SwitchHandler)

	log.Println("Server running on http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
