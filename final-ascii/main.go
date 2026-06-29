package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type pd struct {
	Text   string
	Banner string
	Result string
}

func load(file string) ([]string, error) {
	data, _ := os.ReadFile("banners/" + file + ".txt")
	ln := strings.Split(string(data), "\n")
	return ln, nil
}

func spl(t string) []string {
	t = strings.ReplaceAll(t, "\r\n", "\n")
	return strings.Split(t, "\n")
}

func gen(f []string, w []string) string {
	var res strings.Builder
	for index, wod := range w {
		if wod == "" {
			res.WriteString("\n")
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, r := range wod {
				res.WriteString(f[i+int(r-32)*9])
			}
			res.WriteString("\n")
		}
		if index < len(w)-1 && w[index+1] != "" {
			res.WriteString("\n")
		}
	}
	return res.String()
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func ascii(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	f, _ := load(banner)
	ws := spl(text)
	res := gen(f, ws)
	data := pd{
		Text:   text,
		Banner: banner,
		Result: res,
	}
	tmpl.Execute(w, data)
}

func switcha(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii", ascii)
	http.HandleFunc("/ascii-switch", switcha)
	log.Println("Server up and angry...")
	http.ListenAndServe(":8080", nil)
}
