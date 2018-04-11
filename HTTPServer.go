package describe

import "net/http"

// TypeHTTPServer http server function collections struct
type TypeHTTPServer struct {
	ServeMux   *http.ServeMux
	corsOrigin *TypeStringSlice
	corsMethod *TypeStringSlice
	corsHeader *TypeStringSlice
	Addr       string
	routes     []*TypeHTTPRoute
}

// TypeHTTPRoute http server route function collections struct
type TypeHTTPRoute struct {
	Path         string
	handleMethod map[string]http.HandlerFunc
	// Handler TypeHTTPHandleFunc
}

// TypeHTTPHandlerFunc http server route handler func
type TypeHTTPHandlerFunc func(http.ResponseWriter, *TypeHTTP)

// HTTPServer get *TypeHTTPServer
func HTTPServer() *TypeHTTPServer {
	return &TypeHTTPServer{
		ServeMux:   http.NewServeMux(),
		corsOrigin: StringSlice(nil),
		corsMethod: StringSlice(nil),
		corsHeader: StringSlice(nil),
	}
}

// ListenOn addr that server will listen on
func (hs *TypeHTTPServer) ListenOn(addr string) *TypeHTTPServer {
	hs.Addr = addr
	return hs
}

// AllowOrigin CORS allow origins
func (hs *TypeHTTPServer) AllowOrigin(host string, more ...string) *TypeHTTPServer {
	hs.corsOrigin.Push(host, more...)
	return hs
}

// AllowMethod CORS allow method
func (hs *TypeHTTPServer) AllowMethod(method string, more ...string) *TypeHTTPServer {
	hs.corsMethod.Push(method, more...)
	return hs
}

// AllowHeader CORS allow header
func (hs *TypeHTTPServer) AllowHeader(header string, more ...string) *TypeHTTPServer {
	hs.corsHeader.Push(header, more...)
	return hs
}

// Start start server listening
func (hs *TypeHTTPServer) Start() error {
	for _, hr := range hs.routes {
		hs.ServeMux.HandleFunc(hr.Path, hr.handler)
	}
	return http.ListenAndServe(hs.Addr, hs.ServeMux)
}

// Route get route then config it
func (hs *TypeHTTPServer) Route(path string) *TypeHTTPRoute {
	hr := &TypeHTTPRoute{
		Path:         path,
		handleMethod: map[string]http.HandlerFunc{},
	}
	hs.routes = append(hs.routes, hr)
	return hr
}

func convertHTTPHandlerFunc(f TypeHTTPHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, converNativeRequest(r))
	}
}

// GET handle GET at this route
func (hr *TypeHTTPRoute) GET(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["GET"] = convertHTTPHandlerFunc(f)
	return hr
}

// POST handle POST at this route
func (hr *TypeHTTPRoute) POST(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["POST"] = convertHTTPHandlerFunc(f)
	return hr
}

// PUT handle PUT at this route
func (hr *TypeHTTPRoute) PUT(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["PUT"] = convertHTTPHandlerFunc(f)
	return hr
}

// PATCH handle PATCH at this route
func (hr *TypeHTTPRoute) PATCH(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["PATCH"] = convertHTTPHandlerFunc(f)
	return hr
}

// DELETE handle DELETE at this route
func (hr *TypeHTTPRoute) DELETE(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["DELETE"] = convertHTTPHandlerFunc(f)
	return hr
}

// // OPTIONS handle OPTIONS at this route
// func (hr *TypeHTTPRoute) OPTIONS(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
// 	hr.handleMethod["OPTIONS"] = f
// 	return hr
// }

func (hr *TypeHTTPRoute) handler(w http.ResponseWriter, r *http.Request) {
	f, ok := hr.handleMethod[r.Method]
	if !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		f(w, r)
	}
}
