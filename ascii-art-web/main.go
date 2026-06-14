package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Ascii struct {
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	log.Println("Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	input := r.FormValue("input")
	banner := r.FormValue("banner")

	tmpl, _ := template.ParseFiles("templates/index.html")

	if banner == "" {
		tmpl.Execute(w, Ascii{Error: "Please select a banner"})
		return
	}

	if input == "" {
		tmpl.Execute(w, Ascii{Error: "Please enter some text"})
		return
	}

	font, err := loadBanner(banner)
	if err != nil {
		tmpl.Execute(w, Ascii{Error: "Cannot load banner: " + banner})
		return
	}

	words := splitInput(input)
	result := generateArt(font, words)

	tmpl.Execute(w, Ascii{Result: result})
}

func loadBanner(file string) ([]string, error) {
	file = strings.TrimSuffix(file, ".txt")
	data, err := os.ReadFile("banners/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	return lines, nil
}

func splitInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\\n")
}

func generateArt(fonts []string, words []string) string {
	var res strings.Builder
	for index, word := range words {
		if word == "" {
			res.WriteString("\n")
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, ch := range word {
				idx := i + (int(ch-32) * 9)
				if idx < len(fonts) && idx >= 0 {
					res.WriteString(fonts[idx])
				}
			}
			res.WriteString("\n")
		}
		if index < len(words)-1 && words[index+1] != "" {
			res.WriteString("\n")
		}
	}
	return res.String()
}
