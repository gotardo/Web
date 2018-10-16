package Web

import (
	"log"
	"net/http"
)

func RequestLogger() HandlerComposite {
	return func (h http.HandlerFunc) http.HandlerFunc {
		return func(responseWriter http.ResponseWriter, request *http.Request) {
			log.Print("REQ URL ", request.URL)
			log.Print("REQ HEAD ", request.Header)
			log.Print("REQ BODY ", request.Body)

			h.ServeHTTP(responseWriter, request)

			log.Print("RES STATUS CODE ", responseWriter.Header().Get("statusCode"))
			log.Print("RES HEAD ", responseWriter.Header())
			log.Print("RES BODY ", responseWriter)
		}
	}
}
