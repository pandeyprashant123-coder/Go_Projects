package main

import (
	"encoding/json"
	"net/http"
)

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
	if num := numbers.Number2; num == 0 {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
	}
	sum := numbers.Number1 * numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully Multiplied two numbers",
	}
	json.NewEncoder(w).Encode(response)
}
