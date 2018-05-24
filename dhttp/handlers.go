package dhttp

import "net/http"

// handler for NotFound
func handlerNotFound(w *ResponseWriter, r *Request) {
	w.WriteHeader(http.StatusNotFound)
}
