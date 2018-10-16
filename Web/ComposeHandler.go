package Web

import "net/http"

// Adds a list of middlewares to an http.HandlerFunc.
func ComposeHandler(handler http.HandlerFunc, middlewareList ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewareList {
		handler = middleware(handler)
	}

	return handler
}
