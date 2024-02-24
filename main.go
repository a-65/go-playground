package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Docker contaiter!\nThis docker-image run on server!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}