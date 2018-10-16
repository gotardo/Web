package Web

import "net/http"

// Returns a HandlerComposite function that will evaluate basic auth against BasicAuthCredentials
// before passing the control to the http handler.
func BasicAuth(credentials BasicAuthCredentials) HandlerComposite {
	return func (h http.HandlerFunc) http.HandlerFunc {
		return func(responseWriter http.ResponseWriter, request *http.Request) {
			const UnauthorizedStatusCode = 401
			const UnauthorizedMessage = "Not authorized"

			responseWriter.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			username, password, authOK := request.BasicAuth()

			if !authOK {
				http.Error(responseWriter, UnauthorizedMessage, UnauthorizedStatusCode)
				return;
			}

			if username != credentials.Username || password != credentials.Password {
				http.Error(responseWriter, UnauthorizedMessage, UnauthorizedStatusCode)
				return;
			}

			h.ServeHTTP(responseWriter, request)
		}
	}
}
