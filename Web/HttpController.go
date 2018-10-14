package Web

import "net/http"

type HttpController func(http.ResponseWriter, *http.Request) error
