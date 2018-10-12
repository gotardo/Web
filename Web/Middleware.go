package Web

import "net/http"

// A Middleware function decorates a http.HandlerFunc and is capable of taking decisions about passing the control to
// the handle func or stoping it.
type Middleware func (h http.HandlerFunc) http.HandlerFunc

// Adds a list of middlewares to an http.HandlerFunc.
func ComposeMiddleware(handler http.HandlerFunc, middlewareList ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewareList {
		handler = middleware(handler)
	}

	return handler
}
