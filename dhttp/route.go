package dhttp

import D "github.com/NeoJRotary/describe-go"

// Route http server route function collections struct
type Route struct {
	// Path           string
	Segment        string
	SubRoutes      map[string]*Route
	middlewares    []*Middleware
	methodHandlers map[string]RouteHandlerFunc
}

// RouteHandlerFunc route handler func
type RouteHandlerFunc func(w *ResponseWriter, r *Request)

func newRoute(seg string) *Route {
	return &Route{
		Segment:     seg,
		SubRoutes:   map[string]*Route{},
		middlewares: []*Middleware{},
		methodHandlers: map[string]RouteHandlerFunc{
			"Default": handlerNotFound,
		},
	}
}

// Use use middleware. paths are prefixes for path matching. By default it match all.
func (r *Route) Use(mw *Middleware, paths ...string) *Route {
	mw.paths = paths
	r.middlewares = append(r.middlewares, mw)
	return r
}

// Route get/create route under current route
func (r *Route) Route(path string) *Route {
	segs := D.String(path).Split("/").Get()
	route := r
	for _, seg := range segs {
		route = route.SubSegment(seg)
	}
	return route
}

// GetRoute get subroute by path. Return nil if not found.
func (r *Route) GetRoute(path string) *Route {
	segs := D.String(path).Split("/").Get()
	route := r
	for _, seg := range segs {
		next, ok := r.SubRoutes[seg]
		if !ok {
			return nil
		}
		route = next
	}
	return route
}

// HasRoute check subroute exist by path
func (r *Route) HasRoute(path string) bool {
	return r.GetRoute(path) != nil
}

// SubSegment get/create a segment subroute under current route
func (r *Route) SubSegment(seg string) *Route {
	route, ok := r.SubRoutes[seg]
	if ok {
		return route
	}
	route = newRoute(seg)
	r.SubRoutes[seg] = route
	return route
}

// GetSubSegment get subroute by segment. Return nil if not found.
func (r *Route) GetSubSegment(seg string) *Route {
	route, ok := r.SubRoutes[seg]
	if ok {
		return route
	}
	return nil
}

// HasSubSegment check subroute exist by segment
func (r *Route) HasSubSegment(seg string) bool {
	_, ok := r.SubRoutes[seg]
	return ok
}

// Default default handler
func (r *Route) Default(f RouteHandlerFunc) *Route {
	r.methodHandlers["Default"] = f
	return r
}

// GET handle GET at this route
func (r *Route) GET(f RouteHandlerFunc) *Route {
	r.methodHandlers["GET"] = f
	return r
}

// POST handle POST at this route
func (r *Route) POST(f RouteHandlerFunc) *Route {
	r.methodHandlers["POST"] = f
	return r
}

// PUT handle PUT at this route
func (r *Route) PUT(f RouteHandlerFunc) *Route {
	r.methodHandlers["PUT"] = f
	return r
}

// PATCH handle PATCH at this route
func (r *Route) PATCH(f RouteHandlerFunc) *Route {
	r.methodHandlers["PATCH"] = f
	return r
}

// DELETE handle DELETE at this route
func (r *Route) DELETE(f RouteHandlerFunc) *Route {
	r.methodHandlers["DELETE"] = f
	return r
}

// OPTIONS handle OPTIONS at this route
func (r *Route) OPTIONS(f RouteHandlerFunc) *Route {
	r.methodHandlers["OPTIONS"] = f
	return r
}

// TRACE handle TRACE at this route
func (r *Route) TRACE(f RouteHandlerFunc) *Route {
	r.methodHandlers["TRACE"] = f
	return r
}

// HEAD handle TRACE at this route
func (r *Route) HEAD(f RouteHandlerFunc) *Route {
	r.methodHandlers["HEAD"] = f
	return r
}

// CONNECT handle TRACE at this route
func (r *Route) CONNECT(f RouteHandlerFunc) *Route {
	r.methodHandlers["CONNECT"] = f
	return r
}
