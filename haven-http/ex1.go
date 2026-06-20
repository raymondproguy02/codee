package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func query(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		text := r.URL.Query().Get("name")
		if text == "" {
			text = "Hello, Guest!"
		}
		fmt.Fprintf(w, "Hello, %s!", text)
	} else {
		http.Error(w, "not", 405)
		return
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Send a POST request with text to count words")
	}
	if r.Method == http.MethodPost {
		text, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", 500)
			return
		}
		fmt.Fprintln(w, len(string(text)))
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	con, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	conv, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	if op == "add" {
		fmt.Fprintln(w, con+conv)
	}
	if op == "subtract" {
		fmt.Fprintln(w, con+conv)
	}
	if op == "multiply" {
		fmt.Fprintln(w, con*conv)
	} else {
		http.Error(w, "not an op", 400)
	}
}

func dash(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("X-API-Key")
	if head != "secret123" {
		http.Error(w, "not auth", 401)
		return
	}
	fmt.Fprintln(w, "Welcome")
}

func legacy(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/legacy" {
		http.Redirect(w, r, "http://localhost:5000/v2", http.StatusMovedPermanently)
	}
}

func v2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/hello", query)
	http.HandleFunc("/count", count)
	http.HandleFunc("/calculate", calc)
	http.HandleFunc("/dashboard", dash)
	http.HandleFunc("/legacy", legacy)
	http.HandleFunc("/v2", v2)
	log.Println("App running on http://127.0.0.1:5000/")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
