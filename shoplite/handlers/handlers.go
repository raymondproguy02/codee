package handlers

import (
	"net/http"
	"text/template"
)

type appData struct {
	ID          int
	Name        string
	Item        string
	Price       string
	Description string
}

var Listings []appData
var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("public/index.html")
	if err != nil {
		panic(err)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	item := r.FormValue("item")
	price := r.FormValue("price")
	des := r.FormValue("des")

	if name == "" && item == "" && price == "" && des == "" {
		http.Error(w, "All fields must not be empty", http.StatusBadRequest)
		return
	}

	newData := appData{
		ID:          len(Listings) + 1,
		Name:        name,
		Item:        item,
		Price:       price,
		Description: des,
	}
	Listings = append(Listings, newData)

	data := struct {
		Listings []appData
		NewItem  appData
	}{
		Listings: Listings,
		NewItem:  newData,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendring template", http.StatusInternalServerError)
		return
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := struct {
		Listings []appData
	}{
		Listings: Listings,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendring template", http.StatusInternalServerError)
		return
	}
}
