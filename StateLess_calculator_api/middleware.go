package main

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type Client struct {
	Limiter *rate.Limiter
}

var clients = make(map[string]*Client)
var mu sync.Mutex


func getClientRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if client, exists := clients[ip]; exists {
		return client.Limiter
	}

	limiter := rate.NewLimiter(100, 1)
	clients[ip] = &Client{Limiter: limiter}
	return limiter
}

func rateLimitingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := getClientRateLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

