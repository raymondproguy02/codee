package main

import (
	"fmt"
	"io"
	"net/http"
)

type question struct {
	Text    string
	Options []string
	Answer  int
}

func quest(w http.ResponseWriter, r *http.Request) {
	quiz := []question{
		{
			Text:    "What is the only loop keyword in Go?",
			Options: []string{"while", "for", "loop", "foreach"},
			Answer:  2,
		},
		{
			Text:    "How do you declare a package in Go?",
			Options: []string{"import main", "package main", "pkg main", "define main"},
			Answer:  2,
		},
		{
			Text:    "Which symbol is used for short variable declaration?",
			Options: []string{"=", "==", ":=", "->"},
			Answer:  3,
		},
		{
			Text:    "What is the default value of an uninitialized int variable in Go?",
			Options: []string{"null", "0", "undefined", "-1"},
			Answer:  2,
		},
		{
			Text:    "Which keyword is used to delay the execution of a function until the surrounding function returns?",
			Options: []string{"delay", "later", "defer", "wait"},
			Answer:  3,
		},
		{
			Text:    "How do you start a concurrent goroutine in Go?",
			Options: []string{"go myFunction()", "start myFunction()", "thread myFunction()", "run myFunction()"},
			Answer:  1,
		},
		{
			Text:    "Which built-in function is used to get the number of elements in a slice?",
			Options: []string{"size()", "length()", "count()", "len()"},
			Answer:  4,
		},
		{
			Text:    "What data structure does a Go slice build on top of?",
			Options: []string{"Array", "Linked List", "Map", "Tuple"},
			Answer:  1,
		},
		{
			Text:    "How do you format and print strings without displaying them on the console?",
			Options: []string{"fmt.Printf()", "fmt.Sprintf()", "fmt.Println()", "fmt.Print()"},
			Answer:  2,
		},
		{
			Text:    "What happens when you try to compile a Go program with an imported but unused package?",
			Options: []string{"It compiles with a warning", "It ignores the package automatically", "It fails to compile with an error", "It crashes at runtime"},
			Answer:  3,
		},
	}
}

func utils(w http.ResponseWriter, r *http.Request) {
	gameOver := "Game Over, You can try again."
	fmt.Println(gameOver)

	notValid := "Opps, the answer you provide is invalid."
	fmt.Println(notValid)

	gameWon := "Congradulations %s, you successfully won $100, contact the developers of this site to recive payment"
}

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to the officail Golang Billionaire Game!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root)

	http.ListenAndServe(":8000", mux)
	fmt.Println("Server running on PORT 8000")
}
