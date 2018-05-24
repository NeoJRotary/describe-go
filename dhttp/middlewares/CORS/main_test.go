package cors

import (
	"net/http"
	"testing"

	dhttp "github.com/NeoJRotary/describe-go/dhttp"
)

func newReq() *dhttp.Request {
	r := &http.Request{Header: make(http.Header)}
	return &dhttp.Request{Request: r, MiddlewareValues: nil}
}

func TestCORS_NeedOrigin(t *testing.T) {
	req := newReq()
	header := getHeader(req, Config{NeedOrigin: true})
	if len(header) != 0 {
		t.Error("should skip")
	}

	header = getHeader(req, Config{NeedOrigin: false})
	if len(header) == 0 {
		t.Error("should have header")
	}

	req.Header.Set("Origin", "123")
	header = getHeader(req, Config{NeedOrigin: true})
	if len(header) == 0 {
		t.Error("should have header")
	}
}

func TestCORS_Header(t *testing.T) {
	req := newReq()
	req.Header.Set("Origin", "123")
	req.Header.Set("Access-Control-Request-Headers", "123, 456, 789")

	header := getHeader(req, Config{})
	if header["Access-Control-Allow-Origin"] != "*" {
		t.Error("Access-Control-Allow-Origin should be *")
	}
	if header["Access-Control-Allow-Methods"] != "GET, HEAD, POST" {
		t.Error("Access-Control-Allow-Methods should be GET, HEAD, POST")
	}
	if header["Access-Control-Allow-Headers"] != "123, 456, 789" {
		t.Error("Access-Control-Allow-Headers should be same as Access-Control-Request-Headers")
	}

	header = getHeader(req, Config{Origins: []string{"123"}})
	if header["Access-Control-Allow-Origin"] != "123" {
		t.Error("Access-Control-Allow-Origin should be 123")
	}
	header = getHeader(req, Config{Origins: []string{"456"}})
	if header["Access-Control-Allow-Origin"] != "" {
		t.Error("Access-Control-Allow-Origin should be empty")
	}

	header = getHeader(req, Config{Methods: []string{"GET", "PUT"}})
	if header["Access-Control-Allow-Methods"] != "GET, PUT" {
		t.Error("Access-Control-Allow-Methods should be GET, PUT")
	}

	header = getHeader(req, Config{Headers: []string{"123", "456"}})
	if header["Access-Control-Allow-Headers"] != "123, 456" {
		t.Error("Access-Control-Allow-Headers should be 123, 456")
	}
}
