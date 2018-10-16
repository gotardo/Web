package Web

import "net/http"

// A HandlerComposite function decorates a http.HandlerFunc and is capable of taking decisions about passing the control to
// the handle func or stoping it.
type HandlerComposite func (h http.HandlerFunc) http.HandlerFunc
