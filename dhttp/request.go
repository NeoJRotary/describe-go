package dhttp

import (
	"io/ioutil"
	"net"
	"net/http"

	D "github.com/NeoJRotary/describe-go"
)

// Request inheritance of http.Request
type Request struct {
	*http.Request
	MiddlewareValues map[string]interface{}
}

// ReadAllBody read all body in []byte
func (r *Request) ReadAllBody() []byte {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if D.IsErr(err) {
		return nil
	}
	return b
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
