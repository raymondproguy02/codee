package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

type RateLimiter struct {
    visits map[string]int
    mu     sync.Mutex
    limit  int
    window time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    rl := &RateLimiter{
        visits: make(map[string]int),
        limit:  limit,
        window: window,
    }
    go rl.cleanup()
    return rl
}

func (rl *RateLimiter) cleanup() {
    for {
        time.Sleep(rl.window)
        rl.mu.Lock()
        rl.visits = make(map[string]int)
        rl.mu.Unlock()
    }
}

func (rl *RateLimiter) Middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        rl.mu.Lock()
        rl.visits[ip]++
        count := rl.visits[ip]
        rl.mu.Unlock()

        if count > rl.limit {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

        next(w, r)
    }
}

func main() {
    rl := NewRateLimiter(10, time.Minute)
    
    http.HandleFunc("/", rl.Middleware(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OK - Request %d", rl.visits[r.RemoteAddr])
    }))

    fmt.Println("🚦 Rate limiter running on :8080 (10 req/min)")
    http.ListenAndServe(":8080", nil)
}
