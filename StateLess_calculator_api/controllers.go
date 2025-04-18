package main

import (
	"encoding/json"
	"net/http"
)

type Numbers struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}
type NumArray struct{
	Num []int `json:"num"`
}

type Response struct {
	Result      int 	`json:"result"`
	Description string  `json:"description"`
}

func AddNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	numbers := &Numbers{}
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 + numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully added two numbers",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func SubtractNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	numbers := &Numbers{}
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 - numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully subtracted two numbers",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func MultiplyNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	numbers := &Numbers{}
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum := numbers.Number1 * numbers.Number2
	response := Response{
		Result:      sum,
		Description: "successfully Multiplied two numbers",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func DivideNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	numbers := &Numbers{}
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Sum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	numArray:=&NumArray{}
	err :=json.NewDecoder(r.Body).Decode(&numArray)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	sum:=0
	for _,num:=range numArray.Num{
		sum+=num
	}
	response := Response{
		Result:      sum,
		Description: "successfully added the numbers",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}