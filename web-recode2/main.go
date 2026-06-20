package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

func GenerateArt(fontFileLines []string, words []string) string {
	var result strings.Builder
	for index, word := range words {
		if word == "" {
			continue
		}
		for i := 1; i < 9; i++ {
			for _, char := range word {
				result.WriteString(fontFileLines[i+(int(char-32)*9)])
			}
			if i < 8 {
				result.WriteString("\n")
			}
		}
		if index < len(words)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}

func LoadBanner(fileName string) ([]string, error) {
	fontFile, err := os.ReadFile("banners/" + fileName + ".txt")
	if err != nil {
		return nil, err
	}
	var fontFileLines = strings.Split(strings.ReplaceAll(string(fontFile), "\r\n", "\n"), "\n")
	return fontFileLines, nil
}

func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func ValidateInput(input string) ([]rune, error) {
	odds := []rune{}
	for _, char := range input {
		if char != '\n' && char != '\r' && (char < ' ' || char > '~') {
			odds = append(odds, char)
			continue
		}
	}
	if len(odds) > 0 {
		return odds, fmt.Errorf("unsupported character: %c", odds)
	}
	return nil, nil
}

// Start HTTP server and route requests for ASCII art generation.
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/ascii-switch", asciiShitch)
	log.Println("\033[32mServer running on http://127.0.0.1:8080\033[0m")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Serve the main HTML page for the ASCII art web interface.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}

// Handle POST requests, generate ASCII art, and return the result page.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// if _, err := Validate(text); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	fonts, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	words := SplitInput(text)
	result := GenerateArt(fonts, words)

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func asciiShitch(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// 	return
	// }

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	// if _, err := Validate(text); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	fonts, err := LoadBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	words := SplitInput(text)
	result := GenerateArt(fonts, words)

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
