package dhttp

import D "github.com/NeoJRotary/describe-go"

// MiddlewareFunc middleware handler func
type MiddlewareFunc func(w *ResponseWriter, r *Request, config interface{}) interface{}

// Middleware http server middleware
type Middleware struct {
	Name    string
	Config  interface{}
	Handler MiddlewareFunc
	paths   []string
}

// Rename rename your middleware to avoid value overwriting
func (mw *Middleware) Rename(name string) {
	mw.Name = name
}

// matchPath path is match or not
func (mw *Middleware) matchPath(path string) bool {
	if len(mw.paths) == 0 {
		return true
	}

	for _, prefix := range mw.paths {
		if D.String(path).HasPrefix(prefix) {
			return true
		}
	}

	return false
}
