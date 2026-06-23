package main

import (
	"encoding/json"
	"net/http"
)

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", 405)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" {
		http.Error(w, "text is required", 400)
		return
	}
	if banner == "" {
		http.Error(w, "banner is required", 400)
		return
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "invalid banner", 400)
		return
	}
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "received",
		"text":    text,
		"banner":  banner,
	})
}
