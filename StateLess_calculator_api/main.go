package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	mux := http.NewServeMux()
	h := &calcHandler{}
	mux.Handle("/calc/", rateLimitingMiddleWare(h))
	fmt.Println("Server running on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to calculator app")
}
