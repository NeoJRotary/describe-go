package cors

import (
	"net/http"

	D "github.com/NeoJRotary/describe-go"
	dhttp "github.com/NeoJRotary/describe-go/dhttp"
)

// Config realip middleware config
type Config struct {
	NeedOrigin bool
	Origins    []string
	Methods    []string
	Headers    []string
}

// Middleware to allow CORS
func Middleware(config ...Config) *dhttp.Middleware {
	mw := &dhttp.Middleware{
		Name: "CORS",
		Config: Config{
			NeedOrigin: true,
		},
		Handler: handler,
	}

	if len(config) != 0 {
		mw.Config = config[0]
	}

	return mw
}

// handler Middleware CORS Handler
func handler(w *dhttp.ResponseWriter, r *dhttp.Request, config interface{}) interface{} {
	conf := config.(Config)

	for k, v := range getHeader(r, conf) {
		w.Header().Set(k, v)
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		w.Write(nil)
	}

	return nil
}

func getHeader(r *dhttp.Request, conf Config) map[string]string {
	header := map[string]string{}

	if r.Method != http.MethodOptions && conf.NeedOrigin && r.Header.Get("Origin") == "" {
		return header
	}

	// Allow Origin
	origin := r.Header.Get("Origin")
	if conf.Origins == nil {
		header["Access-Control-Allow-Origin"] = "*"
	} else if D.StringSlice(conf.Origins).Has(origin) {
		header["Access-Control-Allow-Origin"] = origin
	} else {
		header["Access-Control-Allow-Origin"] = ""
	}

	// Allow Methods
	if conf.Methods == nil {
		header["Access-Control-Allow-Methods"] = "GET, HEAD, POST"
	} else {
		header["Access-Control-Allow-Methods"] = D.StringSlice(conf.Methods).Join(", ").Get()
	}

	// Allow Headers
	if conf.Headers == nil {
		header["Access-Control-Allow-Headers"] = r.Header.Get("Access-Control-Request-Headers")
	} else {
		header["Access-Control-Allow-Headers"] = D.StringSlice(conf.Headers).Join(", ").Get()
	}

	return header
}
