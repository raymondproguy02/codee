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
}

func LoadBanner(file string) ([]string, error) {
	file = strings.TrimSuffix(file, ".txt")
	data, err := os.ReadFile("banners/" + file + ".txt")
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Empty banner file")
	}
	var lines = strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	if len(lines) < 855 || len(lines) > 856 {
		return nil, fmt.Errorf("Incomplete banner file")
	}
	return lines, nil
}

func SplitInput(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	return strings.Split(input, "\n")
}

func Validate(input string) (rune, error) {
	for _, r := range input {
		if r != '\r' && r != '\n' && (r < ' ' || r > '~') {
			return r, fmt.Errorf("Invalid character")
		}
	}
	return 0, nil
}

func GenerateArt(fonts []string, words []string) string {
	var res strings.Builder
	for index, word := range words {
		if word == "" {
			res.WriteString("\n")
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, ch := range word {
				res.WriteString(fonts[i+(int(ch-32)*9)])
			}
			res.WriteString("\n")
		}
		if index < len(words)-1 && words[index+1] != "" {
			res.WriteString("\n")
		}
	}
	return res.String()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not allowed", http.StatusNotFound)
		return
	}

	tmlp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "error serving template", http.StatusInternalServerError)
		return
	}
	tmlp.Execute(w, nil)
}
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "not alowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "not alowed", http.StatusNotFound)
		return
	}
	text := r.URL.Query().Get("input")
	banner := r.URL.Query().Get("banner")

	font, _ := LoadBanner(banner)
	word := SplitInput(text)
	res := GenerateArt(font, word)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "error", http.StatusNotFound)
		return
	}
	data := PageData{
		Input:  text,
		Banner: banner,
		Result: res,
	}
	tmpl.Execute(w, data)
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "not alowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "not alowed", http.StatusNotFound)
		return
	}
	text := r.URL.Query().Get("input")
	banner := r.URL.Query().Get("banner")

	if banner == "" {
		r.FormValue("banner")
	}

	font, _ := LoadBanner(banner)
	word := SplitInput(text)
	res := GenerateArt(font, word)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "error", http.StatusNotFound)
		return
	}
	data := PageData{
		Input:  text,
		Banner: banner,
		Result: res,
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ascii-art", asciiHandler)
	mux.HandleFunc("/ascii-switch", switchHandler)

	log.Println("Server Running at http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(err)
	}
}
