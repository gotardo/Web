package Web

import "net/http"

// Handle error is an implementation of the ServeHTTP functio that wille manage any error retorned by an HTTP controller
// and send a 500 status in case of error.
func (innerController HttpController) HandleError(w http.ResponseWriter, r *http.Request) {
	if err := innerController(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
