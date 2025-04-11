package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/add", AddNumbers)
	http.HandleFunc("/subtract", SubtractNumbers)
	http.HandleFunc("/multiply", MultiplyNumbers)
	http.HandleFunc("/divide", DivideNumbers)
	fmt.Println("Server running on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to calculator app")
}
