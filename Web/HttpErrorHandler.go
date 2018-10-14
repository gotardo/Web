package Web

import "net/http"

// Sends a 500 status in case of error.
func (innerController HttpController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := innerController(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
