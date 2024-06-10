package middleware

import (
	"log"
	"net/http"
)

func HeaderWriting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("HeaderWriting")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
