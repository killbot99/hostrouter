package main

import (
	"log"
	"net/http"

	"github.com/killbot99/hostrouter/internal/middleware"
)

func main() {
	http.HandleFunc("/", HandlerPlusMiddlewareChain(
		ForwardingHandler, 
		middleware.Logging,
		middleware.HeaderWriting,
		middleware.Routing,
	))
	http.ListenAndServe(":8080", nil)
}

func ForwardingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Forwarding")
}

// HandlerPlusMiddlewareChain takes an input handler function, and a slice of middlewares functions.
// Then it chains multiple middlwares together in the order they are passed.
// E.G. middlewares[0] is first, middlewares[1] is second, etc...
func HandlerPlusMiddlewareChain(h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h).ServeHTTP
	}
	return h
}
