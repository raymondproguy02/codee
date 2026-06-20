package main

import (
	"fmt"
	"io"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		defer r.Body.Close()
		return
	}

	if len(body) == 0 {
		http.Error(w, "body cannot be empty", 405)
		return
	}
	fmt.Fprintln(w, string(body))
}
