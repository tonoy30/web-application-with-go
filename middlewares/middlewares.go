package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Type alias for middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging: Used to logging
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Printf("[%s]: %s %s\n", r.Method, r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// Method: http method guard
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

// Chain: chaining all the middlewares all together
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
