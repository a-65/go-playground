package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Thenk You, John!!!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}