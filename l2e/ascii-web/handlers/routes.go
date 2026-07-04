package handlers

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Message string `message:"string"`
	Status  string `status:"string"`
}

type AboutPage struct {
	Project     string `json:"project"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Author      string `json:"author"`
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
