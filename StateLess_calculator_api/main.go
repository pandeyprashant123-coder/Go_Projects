package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	h := &calcHandler{}
	mux.Handle("/calc/", rateLimitingMiddleWare(h))
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to calculator app")
}
