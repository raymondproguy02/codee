package handlers

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Message string
	Status  string
}

type AboutPage struct {
	Project     string
	Name        string
	Version     string
	Url         string
	Description string
	Author      string
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	health := Health{
		Message: "Healthy",
		Status:  "OK",
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(health); err != nil {
		http.Error(w, "intern Server error", 500)
		return
	}
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := AboutPage{
		Project:     "01Edu Project",
		Name:        "Ascii Art Web",
		Version:     "2.0.6",
		Url:         "github.com/raymondproguy02/ascii-web",
		Description: "Web version of ascii art!",
		Author:      "raymondproguy",
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		http.Error(w, "intern Server error", 500)
		return
	}
}
