package main

import (
	"net/http"
	"regexp"
)

type calcHandler struct{}
type userHandler struct{}

var (
	add      = regexp.MustCompile(`^\/add[\/]*$`)
	subtract = regexp.MustCompile(`^\/subtract[\/]*$`)
	multiply = regexp.MustCompile(`^\/multiply[\/]*$`)
	divide   = regexp.MustCompile(`^\/divide[\/]*$`)
	sum      = regexp.MustCompile(`^\/sum[\/]*$`)
	login    = regexp.MustCompile(`^\/login[\/]*$`)
	register = regexp.MustCompile(`^\/register[\/]*$`)
)

func (h *calcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && add.MatchString(r.URL.Path):
		h.AddNumbers(w, r)
		return
	case r.Method == http.MethodPost && subtract.MatchString(r.URL.Path):
		h.SubtractNumbers(w, r)
		return
	case r.Method == http.MethodPost && multiply.MatchString(r.URL.Path):
		h.MultiplyNumbers(w, r)
		return
	case r.Method == http.MethodPost && divide.MatchString(r.URL.Path):
		h.DivideNumbers(w, r)
		return
	case r.Method == http.MethodPost && sum.MatchString(r.URL.Path):
		h.Sum(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func(h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	switch{
	case r.Method == http.MethodPost && register.MatchString(r.URL.Path):
		return
	default:
		notFound(w,r)
		return
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
