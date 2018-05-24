package realip

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

func testHandler(req *dhttp.Request, config Config) string {
	req.MiddlewareValues["RealIP"] = getIP(req, config)
	return GetIPV4(req)
}

func TestRealIP(t *testing.T) {
	req := newReq()
	req.Header.Set("X-Forwarded-For", "1.1.1.1, 1.2.3.4, 5.6.7.8")
	req.Header.Set("X-Real-IP", "2.2.2.2")

	config := Config{FromHeader: "MyHead"}
	req.Header.Set("MyHead", "6.6.6.6")
	if testHandler(req, config) != "6.6.6.6" {
		t.Error("Config Error")
	}
	if testHandler(req, Config{}) != "1.1.1.1" {
		t.Error("X-Forwarded-For Error")
	}

	req.Header.Set("X-Forwarded-For", "123123")
	if testHandler(req, Config{}) != "2.2.2.2" {
		t.Error("X-Real-IP Error")
	}

	req.Header.Set("X-Real-IP", "123123")
	if testHandler(req, Config{}) != "127.0.0.1" {
		t.Error("Should get loopback")
	}
}
