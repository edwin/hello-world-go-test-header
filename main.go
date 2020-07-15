package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", get)
	http.ListenAndServe(":8080", nil)
}

func get(w http.ResponseWriter, r *http.Request) {

	// json header default
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Header01", "Satu")
	w.Header().Set("Header02", "DUA")
	w.Header().Set("Header03", "tIGA")
	w.Header().Set("Header04", "EmPaT")
	w.Header().Set("Header05", "limA")

	fmt.Fprint(w, "{\"hello\":\"world\"}")
	return
}