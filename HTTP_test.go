package describe

import (
	"net/http"
	"testing"
	"time"
)

func TestHTTP_Basic(t *testing.T) {
	server := HTTPServer().ListenOn(":12345")
	server.Route("/test").GET(
		func(w http.ResponseWriter, h *TypeHTTP) {
			switch h.GetQuery("qq") {
			case "A":
				w.Write([]byte("getA"))
			case "B":
				w.Write([]byte("getB"))
			default:
				w.Write([]byte("get none"))
			}
		},
	)
	go server.Start()
	time.Sleep(time.Second)
	res, err := HTTP().GET().AtURL(":12345/test").SetQuery("qq", "A").Do()
	if IsErr(err) {
		t.Fatal(err)
	}
	body := string(res.Body)
	if body != "getA" {
		t.Fatal("wrong res, get:", body)
	}
	res, err = HTTP().POST().AtURL(":12345/test").Do()
	if IsErr(err) {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatal("wrong status, get:", res.StatusCode)
	}
}

func TestHTTP_CORS(t *testing.T) {
	server := HTTPServer().ListenOn(":12345").AllowOrigin("localhost").EnableCORS()
	go server.Start()
	time.Sleep(time.Second)
	res, err := HTTP().POST().AtURL(":12345").SetHeader("Origin", "localhost").Do()
	if IsErr(err) {
		t.Fatal(err)
	}
	if res.StatusCode != 404 {
		t.Fatal("CORS shouldn't reject")
	}
	res, err = HTTP().GET().AtURL(":12345").SetHeader("Origin", "google.com").Do()
	if IsErr(err) {
		t.Fatal(err)
	}
	if res.StatusCode != 403 {
		t.Fatal("CORS should reject")
	}
}
