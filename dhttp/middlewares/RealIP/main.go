package realip

import (
	"net"

	D "github.com/NeoJRotary/describe-go"
	dhttp "github.com/NeoJRotary/describe-go/dhttp"
)

// Config realip middleware config
type Config struct {
	FromHeader string
}

// Middleware to get client IP (in string) from header "X-Forwarded-For" or "X-Real-IP". It will return RemoteAdd if headers not found.
func Middleware(config ...Config) *dhttp.Middleware {
	mw := &dhttp.Middleware{
		Name:    "RealIP",
		Config:  nil,
		Handler: handler,
	}

	if len(config) != 0 {
		mw.Config = config[0]
	}

	return mw
}

// handler Middleware RealIP Handler
func handler(w *dhttp.ResponseWriter, r *dhttp.Request, config interface{}) interface{} {
	// Try user config first
	conf, ok := config.(Config)
	if !ok {
		conf = Config{}
	}

	return getIP(r, conf)
}

func getIP(r *dhttp.Request, conf Config) net.IP {
	if conf.FromHeader != "" {
		fromHeader := r.Header.Get(conf.FromHeader)
		if ip := net.ParseIP(fromHeader); ip != nil {
			return ip
		}
	}

	// Try get from X-Forwarded-For
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		xForwardedFor = D.String(xForwardedFor).Split(",").First().TrimSpace().Get()
	}

	// ip := net.ParseIP(xForwardedFor)
	if ip := net.ParseIP(xForwardedFor); ip != nil {
		return ip
	}

	// X-Forwarded-For Failed
	// Try X-Real-IP
	xRealIP := r.Header.Get("X-Real-IP")
	if ip := net.ParseIP(xRealIP); ip != nil {
		return ip
	}

	return r.GetRemoteIP()
}

// Get get net.IP from request. Return nil if not found.
func Get(r *dhttp.Request) net.IP {
	return GetBy("RealIP", r)
}

// GetBy get net.IP from request by name. Return nil if not found.
func GetBy(name string, r *dhttp.Request) net.IP {
	v, ok := r.MiddlewareValues[name]
	if !ok {
		return nil
	}
	ip, ok := v.(net.IP)
	if !ok {
		return nil
	}
	return ip
}

// GetIPV4 get ipv4 string from request. Return "127.0.0.1" if it is loopback. Return empty string if not found.
func GetIPV4(r *dhttp.Request) string {
	ip := Get(r)
	if ip == nil {
		return ""
	}
	if ip.IsLoopback() {
		return "127.0.0.1"
	}
	return ip.To4().String()
}

// GetIPV6 get ipv6 string from request. Return "::1" if it is loopback. Return empty string if not found.
func GetIPV6(r *dhttp.Request) string {
	ip := Get(r)
	if ip == nil {
		return ""
	}
	if ip.IsLoopback() {
		return "::1"
	}
	return ip.To16().String()
}
