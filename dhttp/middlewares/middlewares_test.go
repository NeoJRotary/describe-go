package middlewares

import (
	"testing"
	"time"

	D "github.com/NeoJRotary/describe-go"
	dhttp "github.com/NeoJRotary/describe-go/dhttp"
	Access "github.com/NeoJRotary/describe-go/dhttp/middlewares/Access"
	Auth "github.com/NeoJRotary/describe-go/dhttp/middlewares/Auth"
	CORS "github.com/NeoJRotary/describe-go/dhttp/middlewares/CORS"
	RealIP "github.com/NeoJRotary/describe-go/dhttp/middlewares/RealIP"
)

func authValidater(w *dhttp.ResponseWriter, r *dhttp.Request) bool {
	return r.Header.Get("token") == "12345"
}

func TestMiddlewares(t *testing.T) {
	defer D.RecoverErr(func(e error) { t.Fatal(e) })

	server := dhttp.Server().ListenOn("localhost:12345").Use(CORS.Middleware()).Use(RealIP.Middleware())
	server.Use(&dhttp.Middleware{
		Handler: func(w *dhttp.ResponseWriter, r *dhttp.Request, config interface{}) interface{} {
			w.Header().Set("CustomMiddleware", "yoyo")
			return nil
		},
	})
	server.Use(Auth.Middleware(true, authValidater), "/deny")

	server.Route("/allow/loopback").Use(Access.Middleware(Access.Config{
		Rules: []string{
			Access.Allow("loopback"),
			Access.Deny("all"),
		},
	}))

	conf := Access.Config{}
	conf.Deny("loopback")
	server.Route("/deny/loopback").Use(Access.Middleware(conf))

	conf = Access.Config{}
	conf.Deny("all")
	server.Route("/deny/all").Use(Access.Middleware(conf))

	server.Route("/deny/ip-range/remote").Use(Access.Middleware(Access.Config{
		Rules: []string{
			Access.Deny("30.0.0.0/8"),
		},
	}))

	server.Route("/deny/ip-range/real").Use(Access.Middleware(Access.Config{
		UseRealIP: true,
		Rules: []string{
			Access.Deny("30.0.0.0/8"),
		},
	}))

	go server.Start()

	client := dhttp.Client().SetTimeout(time.Second)

	res, err := client.HEAD().SetURL("http://localhost:12345").Do()
	D.CheckErr(err)
	if res.Header.Get("CustomMiddleware") != "yoyo" {
		t.Fatal("CustomMiddleware not work")
	}

	res, err = client.GET().SetURL("http://localhost:12345/allow/loopback").Do()
	D.CheckErr(err)
	if res.StatusCode == 403 {
		t.Fatal("should allow loopback")
	}

	res, err = client.GET().SetURL("http://localhost:12345/deny/loopback").Do()
	D.CheckErr(err)
	if res.StatusCode != 401 {
		t.Fatal("should unauth")
	}

	client.SetHeader("token", "12345")
	res, err = client.GET().SetURL("http://localhost:12345/deny/loopback").Do()
	D.CheckErr(err)
	if res.StatusCode == 401 {
		t.Fatal("should not unauth")
	}
	if res.StatusCode != 403 {
		t.Fatal("should deny loopback")
	}

	res, err = client.GET().SetURL("http://localhost:12345/deny/all").Do()
	D.CheckErr(err)
	if res.StatusCode != 403 {
		t.Fatal("should deny")
	}

	res, err = client.GET().SetURL("http://localhost:12345/deny/ip-range/remote").Do()
	D.CheckErr(err)
	if res.StatusCode == 403 {
		t.Fatal("should allow")
	}

	c := client.GET().SetURL("http://localhost:12345/deny/ip-range/real")
	c.SetHeader("X-Forwarded-For", "30.1.1.1, 1.2.3.4, 2.3.4.5")
	res, err = c.Do()
	D.CheckErr(err)
	if res.StatusCode != 403 {
		t.Fatal("should deny")
	}

	c.SetHeader("X-Forwarded-For", "")
	c.SetHeader("X-Real-IP", "30.2.2.2")
	res, err = c.Do()
	D.CheckErr(err)
	if res.StatusCode != 403 {
		t.Fatal("should deny")
	}
}
