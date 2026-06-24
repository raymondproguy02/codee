package main

import (
	"html/template"
	"log"
	"net/http"
	//"strconv"
)

type Calculator struct {
	Num1   string
	Num2   string
	Optn   string
	Result string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", 500)
		return
	}
	tmpl.Execute(w, nil)
}

func calc(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// a := r.FormValue("num1")
	// b := r.FormValue("num2")
	// o := r.FormValue("op")

	// con, err := strconv.Atoi(a)
	// if err != nil {
	// 	http.Error(w, "Must be valid intiger", 400)
	// 	return
	// }
	// conv, err := strconv.Atoi(b)
	// if err != nil {
	// 	http.Error(w, "Must be valid intiger", 400)
	// 	return
	// }
	// if o == "add" {
	// 	//con + conv
	// }
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)

	log.Println("Server running on http://127.0.0.1:8000/")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
