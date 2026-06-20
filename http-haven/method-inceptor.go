package main

import (
	"fmt"
	"net/http"
)

func method(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "You made %s request", r.Method)
	}
	fmt.Fprintf(w, "You made %s request", r.Method)
}
