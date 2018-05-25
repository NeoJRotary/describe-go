package dhttp

import (
	"net"
	"net/http"
)

// Request inheritance of http.Request
type Request struct {
	*http.Request
	MiddlewareValues map[string]interface{}
}

// GetRemoteIP get net.IP from Request.RemoteAddr
func (r *Request) GetRemoteIP() net.IP {
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil
	}
	return net.ParseIP(remoteIP)
}

// SetMiddlewareValue set middleware value
func (r *Request) SetMiddlewareValue(name string, value interface{}) {
	r.MiddlewareValues[name] = value
}
