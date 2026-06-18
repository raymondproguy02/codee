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

func AppHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	item := r.FormValue("item")
	price := r.FormValue("price")
	des := r.FormValue("des")

	idCount := len(id) + 1

	if name == "" && item == "" && price == "" && des == "" {
		http.Error(w, "All fields must not be empty", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "Frontend page missing", http.StatusNotFound)
		return
	}

	data := appData{
		ID:          idCount,
		Name:        name,
		Item:        item,
		Price:       price,
		Description: des,
	}
	temp.Execute(w, data)
	Listings = append(Listings, data)
}
