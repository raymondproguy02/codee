package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	// res, _ := io.ReadAll(data.Body)
	// fmt.Fprintln(w, string(res))
	var artists []Artist
	err = json.NewDecoder(data.Body).Decode(&artists)
	mj, _ := json.MarshalIndent(artists, "", "  ")
	fmt.Fprintln(w, string(mj))
	//fmt.Printf("%+v", artists)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server runnig...")
	http.ListenAndServe(":8080", nil)
}
