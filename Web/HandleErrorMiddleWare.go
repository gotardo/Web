package Web

import "net/http"

// Handle error is an implementation of the ServeHTTP functio that wille manage any error retorned by an HTTP controller
// and send a 500 status in case of error.
type HandleErrorMiddleWare struct {
	inner HttpController
}

func (middleware HandleErrorMiddleWare) Action(w http.ResponseWriter, r *http.Request) {
	if err := middleware.inner.Action(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
