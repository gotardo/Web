package Web

import "net/http"

type HttpController interface {
	Action(http.ResponseWriter, *http.Request) error
}
