package auth

import (
	"net/http"
	"testing"

	dhttp "github.com/NeoJRotary/describe-go/dhttp"
)

func newReq() *dhttp.Request {
	r := &http.Request{
		RemoteAddr: "[::1]:50354",
		Header:     make(http.Header),
	}
	return &dhttp.Request{Request: r, MiddlewareValues: map[string]interface{}{}}
}

func TestRealIP(t *testing.T) {
	req := newReq()
	conf := Config{
		AutoReject: false,
		Validater: func(w *dhttp.ResponseWriter, r *dhttp.Request) bool {
			return r.Header.Get("token") == "1234567890"
		},
	}

	req.Header.Set("token", "1234567890")
	if !handler(nil, req, conf).(bool) {
		t.Fatal("should pass")
	}

	req.Header.Set("token", "wrongggg")
	if handler(nil, req, conf).(bool) {
		t.Fatal("should not pass")
	}
}
