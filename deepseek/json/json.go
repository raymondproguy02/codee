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
}
