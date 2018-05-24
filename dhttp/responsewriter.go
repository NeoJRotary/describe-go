package dhttp

import "net/http"

// ResponseWriter wrapper of http.ResponseWriter
type ResponseWriter struct {
	StatusCode int
	Done       bool
	w          http.ResponseWriter
}

// Write http.ResponseWriter Write
func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.Done = true
	return rw.w.Write(b)
}

// WriteText use string to Write
func (rw *ResponseWriter) WriteText(s string) (int, error) {
	return rw.Write([]byte(s))
}

// WriteHeader http.ResponseWriter WriteHeader
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.w.WriteHeader(statusCode)
	rw.StatusCode = statusCode
}

// Header http.ResponseWriter Header
func (rw *ResponseWriter) Header() http.Header {
	return rw.w.Header()
}

func newResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		StatusCode: 200,
		Done:       false,
		w:          w,
	}
}
