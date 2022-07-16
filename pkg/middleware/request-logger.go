package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		method := r.Method

		log.Printf("Method: %s Path: %s\n", method, url)
		next.ServeHTTP(w, r)
	})

}

func DummyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Hello from dummy middleware")
		next.ServeHTTP(w, r)
	})
}
