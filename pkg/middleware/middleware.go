package middleware

import "net/http"

// Middleware is a type that represents a middleware signiture
type Middleware func(http.Handler) http.Handler

// Chain is a slice of middlewares
type Chain []Middleware

// Then is a method on "Chain" type and allows you to pass the handler that you want to apply middlewares
func (c Chain) Then(handler http.Handler) http.Handler {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	//now we loop them
	for i := range c {
		//we need to apply the last middleware first and come up to the first one and each call to middleare return a new handler
		handler = c[len(c)-1-i](handler)
	}

	return handler
}

//CreateChain is a function that takes as many middleware and applies them to the handler
func CreateChain(middlewares ...Middleware) Chain {
	var slice Chain
	return append(slice, middlewares...)
}
