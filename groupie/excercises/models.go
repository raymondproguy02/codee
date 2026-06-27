package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type Artist struct {
// 	ID           int      `json:"id"`
// 	Image        string   `json:"image"`
// 	Name         string   `json:"name"`
// 	Members      []string `json:"members"`
// 	CreationDate int      `json:"creationDate"`
// 	FirstAlbum   string   `json:"firstAlbum"`
// 	Locations    string   `json:"loctions"`
// 	ConcertDates string   `json:"concertDates"`
// 	Reletions    string   `json:"relations"`
// }

// func h2(w http.ResponseWriter, r *http.Request) {
// 	var artists Artist
// 	fmt.Printf("%+v", artists)
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Server runnig...")
// 	http.ListenAndServe(":8080", nil)
// }
