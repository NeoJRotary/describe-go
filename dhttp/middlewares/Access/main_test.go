package access

import (
	"net"
	"net/http"
	"testing"

	dhttp "github.com/NeoJRotary/describe-go/dhttp"
)

func newReq() *dhttp.Request {
	r := &http.Request{RemoteAddr: "1.1.1.1:1234"}
	return &dhttp.Request{Request: r, MiddlewareValues: map[string]interface{}{}}
}

func TestAccess_GetIP(t *testing.T) {
	r := newReq()

	if getIP(r, Config{}).String() != "1.1.1.1" {
		t.Error("should get RemoteAddr")
	}

	if getIP(r, Config{UseRealIP: true}).String() != "1.1.1.1" {
		t.Error("should get RemoteAddr")
	}

	r.SetMiddlewareValue("RealIP", net.ParseIP("2.2.2.2"))
	if getIP(r, Config{UseRealIP: true}).String() != "2.2.2.2" {
		t.Error("should get RealIP")
	}
}

func testValidate(ip string, c Config) bool {
	r := newReq()
	r.RemoteAddr = ip + ":1234"
	return validate(getIP(r, c), c)
}

func TestAccess_Validate(t *testing.T) {
	c := Config{
		Rules: []string{
			Allow("1.1.1.1"),
			Deny("all"),
		},
	}
	if !testValidate("1.1.1.1", c) {
		t.Error("should allow")
	}
	if testValidate("2.2.2.2", c) {
		t.Error("should deny")
	}

	c = Config{
		Rules: []string{
			Deny("loopback"),
		},
	}
	if !testValidate("1.1.1.1", c) {
		t.Error("should allow")
	}
	if testValidate("127.0.0.1", c) {
		t.Error("should deny")
	}

	c = Config{
		Rules: []string{
			Allow("1.1.1.1"),
			Deny("1.0.0.0/8"),
			Allow("2.2.0.0/16"),
			Deny("2.0.0.0/8"),
		},
	}
	if testValidate("1.1.1.2", c) {
		t.Error("should deny")
	}
	if !testValidate("2.2.2.2", c) {
		t.Error("should allow")
	}
	if testValidate("2.0.2.2", c) {
		t.Error("should deny")
	}
}
