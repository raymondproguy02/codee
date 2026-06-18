package handlers

import (
	"net/http"
	"text/template"
)

type appData struct {
	ID          int
	Name        string
	Item        string
	Price       float32
	Description string
}

var Listings []appData

func AppHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	item := r.FormValue("item")
	price := r.FormValue("price")
	des := r.FormValue("des")

	temp, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "Frontend page missing", http.StatusNotFound)
		return
	}

	data := appData{
		ID:          1,
		Name:        "Ray",
		Item:        "Laptop",
		Price:       240.00,
		Description: "Hp Elite, 8th Gen, core i5",
	}
	temp.Execute(w, data)
	Listings = append(Listings, data)
}
