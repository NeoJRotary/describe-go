package access

import (
	"net"

	D "github.com/NeoJRotary/describe-go"
	dhttp "github.com/NeoJRotary/describe-go/dhttp"
	realip "github.com/NeoJRotary/describe-go/dhttp/middlewares/RealIP"
)

// Config access middleware config
type Config struct {
	UseRealIP bool
	Rules     []string
}

// Allow add allow rule
func (c *Config) Allow(s string) {
	if c.Rules == nil {
		c.Rules = []string{}
	}
	c.Rules = append(c.Rules, "allow "+s)
}

// Deny add deny rule
func (c *Config) Deny(s string) {
	if c.Rules == nil {
		c.Rules = []string{}
	}
	c.Rules = append(c.Rules, "deny "+s)
}

// Allow get rule string of allowance
func Allow(s string) string {
	return "allow " + s
}

// Deny get rule string of denial
func Deny(s string) string {
	return "deny " + s
}

// Middleware access request by allow or deny IP/CIDR
func Middleware(conf Config) *dhttp.Middleware {
	mw := &dhttp.Middleware{
		Name:    "Access",
		Config:  conf,
		Handler: handler,
	}

	return mw
}

// handler Middleware Access Handler
func handler(w *dhttp.ResponseWriter, r *dhttp.Request, config interface{}) interface{} {
	conf := config.(Config)

	if !validate(getIP(r, conf), conf) {
		w.WriteHeader(403)
		w.WriteText("Forbidden IP")
	}

	return nil
}

func getIP(r *dhttp.Request, conf Config) net.IP {
	if conf.UseRealIP {
		if real := realip.Get(r); real != nil {
			return real
		}
	}

	return r.GetRemoteIP()
}

func validate(ip net.IP, conf Config) bool {
	for _, rule := range conf.Rules {
		rs := D.String(rule)

		if rs.HasPrefix("allow ") {
			rule = rs.TrimPrefix("allow ").Get()
			if rule == "all" {
				return true
			}
			if ip == nil {
				continue
			}
			if rule == "loopback" {
				if ip.IsLoopback() {
					return true
				}
			}
			if ruleIP := net.ParseIP(rule); ruleIP != nil {
				if ip.Equal(ruleIP) {
					return true
				}
			}
			_, ipv4net, err := net.ParseCIDR(rule)
			if D.IsErr(err) {
				continue
			}
			if ipv4net.Contains(ip) {
				return true
			}
		}

		if rs.HasPrefix("deny ") {
			rule = rs.TrimPrefix("deny ").Get()
			if rule == "all" {
				return false
			}
			if ip == nil {
				continue
			}
			if rule == "loopback" {
				if ip.IsLoopback() {
					return false
				}
			}
			if ruleIP := net.ParseIP(rule); ruleIP != nil {
				if ip.Equal(ruleIP) {
					return false
				}
			}
			_, ipv4net, err := net.ParseCIDR(rule)
			if D.IsErr(err) {
				continue
			}
			if ipv4net.Contains(ip) {
				return false
			}
		}
	}

	return true
}
