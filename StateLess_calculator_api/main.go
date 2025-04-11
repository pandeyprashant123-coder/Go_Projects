package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numbers struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type Response struct {
	Result      int
	Description string
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/add", AddNumbers)
	http.HandleFunc("/subtract", SubtractNumbers)
	http.HandleFunc("/multiply", MultiplyNumbers)
	fmt.Println("Server running on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to calculator app")
}

func AddNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 + numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully added two numbers",
	}
	json.NewEncoder(w).Encode(response)
}
func SubtractNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 - numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully subtracted two numbers",
	}
	json.NewEncoder(w).Encode(response)
}
func MultiplyNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 * numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully Multiplied two numbers",
	}
	json.NewEncoder(w).Encode(response)
}
func DivideNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 * numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully Multiplied two numbers",
	}
	json.NewEncoder(w).Encode(response)
}
