package dhttp

import (
	"testing"
	"time"

	D "github.com/NeoJRotary/describe-go"
)

func TestDHTTP_Basic(t *testing.T) {
	server := Server().ListenOn("localhost:10001")
	server.Route("/test").GET(func(w *ResponseWriter, r *Request) {
		w.WriteText("123")
	})

	go server.Start()

	res, err := Client().GET().SetURL("http://localhost:10001/test").SetTimeout(time.Second).Do()
	if D.IsErr(err) {
		t.Fatal(err)
	}

	if string(res.ReadAllBody()) != "123" {
		t.Fatal("should get 123")
	}
}

func TestDHTTP_Route(t *testing.T) {
	defer D.RecoverErr(func(e error) { t.Fatal(e) })

	server := Server().ListenOn("localhost:10002")
	server.Route("/test").GET(func(w *ResponseWriter, r *Request) {
		w.WriteText("test")
	})
	server.Route("/test/more/longer").GET(func(w *ResponseWriter, r *Request) {
		w.WriteText("longer")
	})
	server.Route("yoyo").Route("more/longer").SubSegment("with").SubSegment("two").SubSegment("seg").GET(func(w *ResponseWriter, r *Request) {
		w.WriteText("withseg")
	})
	go server.Start()

	c := Client().GET().SetTimeout(time.Second)

	res, err := c.SetURL("http://localhost:10002/test").Do()
	D.CheckErr(err)
	if string(res.ReadAllBody()) != "test" {
		t.Fatal("should get test")
	}

	res, err = c.SetURL("http://localhost:10002/test/more/longer").Do()
	D.CheckErr(err)
	if string(res.ReadAllBody()) != "longer" {
		t.Fatal("should get longer")
	}

	res, err = c.SetURL("http://localhost:10002/yoyo/more/longer/with/two/seg").Do()
	D.CheckErr(err)
	if string(res.ReadAllBody()) != "withseg" {
		t.Fatal("should get withseg")
	}

	res, err = c.SetURL("http://localhost:10002/yoyo/something/worng").Do()
	D.CheckErr(err)
	if res.StatusCode != 404 {
		t.Fatal("route shouldnt found")
	}
}

func TestDHTTP_Method(t *testing.T) {
	defer D.RecoverErr(func(e error) { t.Fatal(e) })

	server := Server().ListenOn("localhost:10003")
	route := server.Route("/test")
	route.GET(func(w *ResponseWriter, r *Request) {
		w.WriteText("get")
	})
	route.POST(func(w *ResponseWriter, r *Request) {
		w.WriteText("post")
	})
	go server.Start()

	c := Client().SetURL("http://localhost:10003/test").SetTimeout(time.Second)

	res, err := c.GET().Do()
	D.CheckErr(err)
	if string(res.ReadAllBody()) != "get" {
		t.Fatal("should get get")
	}

	res, err = c.POST().Do()
	D.CheckErr(err)
	if string(res.ReadAllBody()) != "post" {
		t.Fatal("should get post")
	}

	res, err = c.PUT().Do()
	D.CheckErr(err)
	if res.StatusCode != 404 {
		t.Fatal("should be 404")
	}

}
