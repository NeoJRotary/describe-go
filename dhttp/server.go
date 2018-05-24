package dhttp

import (
	"net/http"

	D "github.com/NeoJRotary/describe-go"
)

// TypeServer http server function collections struct
type TypeServer struct {
	ServeMux    *http.ServeMux
	Addr        string
	middlewares []*Middleware
	rootRoute   *Route
}

// Server create new http server
func Server() *TypeServer {
	return &TypeServer{
		ServeMux:    http.NewServeMux(),
		middlewares: []*Middleware{},
		rootRoute:   newRoute("/"),
	}
}

// ListenOn addr that server will listen on
func (svr *TypeServer) ListenOn(addr string) *TypeServer {
	svr.Addr = addr
	return svr
}

// Use use middleware. paths are prefixes for path matching. By default it match all.
func (svr *TypeServer) Use(mw *Middleware, paths ...string) *TypeServer {
	mw.paths = paths
	svr.middlewares = append(svr.middlewares, mw)
	return svr
}

// Route get route then config it
func (svr *TypeServer) Route(path string) *Route {
	if path == "/" {
		return svr.rootRoute
	}

	return svr.rootRoute.Route(D.String(path).TrimPrefix("/").Get())
}

// Start start server listening
func (svr *TypeServer) Start() error {
	svr.ServeMux.HandleFunc("/", svr.mainHandler)
	return http.ListenAndServe(svr.Addr, svr.ServeMux)
}

// handle middlewares and routes
func (svr *TypeServer) mainHandler(w http.ResponseWriter, r *http.Request) {
	newW := newResponseWriter(w)
	newR := &Request{r, map[string]interface{}{}}

	// go through server middleware
	for _, middleware := range svr.middlewares {
		if !middleware.matchPath(r.URL.Path) {
			continue
		}

		newR.MiddlewareValues[middleware.Name] = middleware.Handler(newW, newR, middleware.Config)
		if newW.Done {
			return
		}
	}

	// search route
	route := svr.rootRoute
	if r.URL.Path != "/" {
		segs := D.String(r.URL.Path).TrimPrefix("/").Split("/").Get()
		for _, seg := range segs {
			next := route.GetSubSegment(seg)
			if next == nil {
				break
			}
			route = next
		}
	}

	// run middlewares
	for _, middleware := range route.middlewares {
		if !middleware.matchPath(r.URL.Path) {
			continue
		}

		newR.MiddlewareValues[middleware.Name] = middleware.Handler(newW, newR, middleware.Config)
		if newW.Done {
			return
		}
	}

	// get handler
	handler, ok := route.methodHandlers[r.Method]
	if !ok {
		handler = route.methodHandlers["Default"]
	}

	// run handler
	if handler != nil {
		handler(newW, newR)
	}
}
