package main

import (
	//"code/perser"

	"fmt"
	"log"
	"net/http"
	//"os"
)

func main() {
	// if len(os.Args) != 3 {
	// 	return
	// }
	// if len(os.Args[1]) == 0 {
	// 	return
	// }
	// valid, err := perser.Validate(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("invalid %q", valid)
	// 	return
	// }
	// font, err := perser.LoadBanner(os.Args[2])
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// word := perser.SplitInput(os.Args[1])
	// res := perser.GenerateArt(font, word)
	// fmt.Println(res)
	//port := "8000"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ascii Art Go Web Server")
	})
	log.Println("Server up and running: Visit http://127.0.0.1:8000/")
	http.ListenAndServe("8080", mux)
}
