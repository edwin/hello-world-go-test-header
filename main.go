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
	w.Header().Set("HeaDer01", "Satu")
	w.Header().Set("HeadeR02", "DUA")
	w.Header().Set("HeAder03", "tIGA")
	w.Header().Set("HEader04", "EmPaT")
	w.Header().Set("HEAdER05", "limA")

	fmt.Fprint(w, "{\"hello\":\"world\"}")
	return
}
