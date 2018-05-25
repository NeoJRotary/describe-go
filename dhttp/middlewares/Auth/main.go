package auth

import (
	dhttp "github.com/NeoJRotary/describe-go/dhttp"
)

// Config auth middleware config
type Config struct {
	AutoReject bool
	Validator  Validator
}

// Validator do auth
type Validator func(w *dhttp.ResponseWriter, r *dhttp.Request) bool

// Middleware to do authentication before handling request
func Middleware(autoReject bool, v Validator) *dhttp.Middleware {
	mw := &dhttp.Middleware{
		Name: "Auth",
		Config: Config{
			AutoReject: autoReject,
			Validator:  v,
		},
		Handler: handler,
	}

	return mw
}

// handler Middleware Auth Handler
func handler(w *dhttp.ResponseWriter, r *dhttp.Request, config interface{}) interface{} {
	conf := config.(Config)
	result := conf.Validator(w, r)

	if conf.AutoReject && !result {
		w.WriteHeader(401)
		w.WriteText("")
	}

	return result
}

// Valid check request is valid or not
func Valid(r *dhttp.Request) bool {
	return ValidBy("Auth", r)
}

// ValidBy check request is valid or not by name
func ValidBy(name string, r *dhttp.Request) bool {
	result, ok := r.MiddlewareValues[name]
	if !ok {
		return true
	}
	v, ok := result.(bool)
	if !ok {
		return true
	}
	return v
}
