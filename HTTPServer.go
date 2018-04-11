package describe

import "net/http"

// TypeHTTPServer http server function collections struct
type TypeHTTPServer struct {
	ServeMux    *http.ServeMux
	corsHandler TypeHTTPCORSHandler
	corsOrigin  *TypeStringSlice
	corsMethod  *TypeStringSlice
	corsHeader  *TypeStringSlice
	corsEnable  bool
	Addr        string
	routes      []*TypeHTTPRoute
}

// TypeHTTPCORSHandler http CORS handler
type TypeHTTPCORSHandler func(defaultResult bool, r *http.Request) bool

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

// AllowOrigin server level CORS allow origins
func (hs *TypeHTTPServer) AllowOrigin(host string, more ...string) *TypeHTTPServer {
	hs.corsOrigin.Push(host, more...)
	return hs
}

// AllowMethod server level CORS allow method
func (hs *TypeHTTPServer) AllowMethod(method string, more ...string) *TypeHTTPServer {
	hs.corsMethod.Push(method, more...)
	return hs
}

// AllowHeader server level CORS allow header
func (hs *TypeHTTPServer) AllowHeader(header string, more ...string) *TypeHTTPServer {
	hs.corsHeader.Push(header, more...)
	return hs
}

// EnableCORS enable server level CORS response
func (hs *TypeHTTPServer) EnableCORS() *TypeHTTPServer {
	hs.corsEnable = true
	return hs
}

// SetCORSHandler setup your CORS handler
func (hs *TypeHTTPServer) SetCORSHandler(f TypeHTTPCORSHandler) *TypeHTTPServer {
	hs.corsHandler = f
	return hs
}

func (hs *TypeHTTPServer) corsCheck(r *http.Request) bool {
	result := true
	result = result && hs.corsOrigin.Has(r.Header.Get("Origin"))

	if hs.corsHandler != nil {
		return hs.corsHandler(result, r)
	}
	return result
}

func (hs *TypeHTTPServer) handle(w http.ResponseWriter, r *http.Request) {
	if hs.corsEnable {
		if !hs.corsMethod.Empty() {
			w.Header().Set("Access-Control-Allow-Methods", hs.corsMethod.Join(",").Get())
		}
		if !hs.corsHeader.Empty() {
			w.Header().Set("Access-Control-Allow-Headers", hs.corsHeader.Join(",").Get())
		}

		if hs.corsCheck(r) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "")
			if r.Method != http.MethodOptions {
				w.WriteHeader(http.StatusForbidden)
			}
			return
		}
	}

	for _, hr := range hs.routes {
		if hr.Path == r.URL.Path {
			f, ok := hr.handleMethod[r.Method]
			if !ok {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			f(w, r)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Start start server listening
func (hs *TypeHTTPServer) Start() error {
	hs.ServeMux.HandleFunc("/", hs.handle)
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

// convertHTTPHandlerFunc
func (hr *TypeHTTPRoute) convertHTTPHandlerFunc(f TypeHTTPHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := converRequestToTypeHTTP(r)
		h.RoutePath = hr.Path
		f(w, h)
	}
}

// GET handle GET at this route
func (hr *TypeHTTPRoute) GET(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["GET"] = hr.convertHTTPHandlerFunc(f)
	return hr
}

// POST handle POST at this route
func (hr *TypeHTTPRoute) POST(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["POST"] = hr.convertHTTPHandlerFunc(f)
	return hr
}

// PUT handle PUT at this route
func (hr *TypeHTTPRoute) PUT(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["PUT"] = hr.convertHTTPHandlerFunc(f)
	return hr
}

// PATCH handle PATCH at this route
func (hr *TypeHTTPRoute) PATCH(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["PATCH"] = hr.convertHTTPHandlerFunc(f)
	return hr
}

// DELETE handle DELETE at this route
func (hr *TypeHTTPRoute) DELETE(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
	hr.handleMethod["DELETE"] = hr.convertHTTPHandlerFunc(f)
	return hr
}

// // OPTIONS handle OPTIONS at this route
// func (hr *TypeHTTPRoute) OPTIONS(f TypeHTTPHandlerFunc) *TypeHTTPRoute {
// 	hr.handleMethod["OPTIONS"] = f
// 	return hr
// }
