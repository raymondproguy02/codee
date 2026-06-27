package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	res, _ := io.ReadAll(data.Body)
	fmt.Println(string(res))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server runnig...")
	http.ListenAndServe(":8080", nil)
}
