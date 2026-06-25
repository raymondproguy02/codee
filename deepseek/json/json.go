package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ProGuy struct {
	Name     string   `json:"name"`
	Bio      string   `json:"bio"`
	Stack    []string `json:"stack"`
	Location string   `json:"location"`
}

func main() {
	// Encoding Go Struct to JSON
	p := ProGuy{
		Name:     "Raymond",
		Bio:      "Founder Crydensync",
		Stack:    []string{"Go", "Git", "Js", "Bash"},
		Location: "Remote",
	}
	jdata, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:", string(jdata))

	// Decoding JSON struct to Go struct
	jstr := `{
		"name": "Raymond", 
		"bio": "Founder Crydensync",
		"stack": ["Go", "Git", "Js", "Bash"],
		"location": "Remote"
	}`
	var person ProGuy
	err = json.Unmarshal([]byte(jstr), &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded: %+v\n", person)
}
