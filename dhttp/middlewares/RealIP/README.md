# Middleware RealIP
DHTTP server middleware to get client real IP from header   
   
## Intro
It will try `X-Forwarded-For` first, than `X-Real-IP`. If both are invalid it parse by `http.Request.RemoteAddr`.

## Config
- FromHeader : try to get from this header first. Which you can use custom header or `X-Real-IP`.

## Method
- Get `(r *DHTTP.Request)` `net.IP`    
  get real IP in `net.IP` from request. Return nil if not found.
- GetBy `(name string, r *DHTTP.Request)` `net.IP`    
  get real IP in `net.IP` from request by name. Return nil if not found.
- GetIPV4 `(r *DHTTP.Request)` `string`    
  get real IP in ipv4. "127.0.0.1" if loopback. Empty string if not found.
- GetIPV6 `(r *DHTTP.Request)` `string`    
  get real IP in ipv6. "::1" if loopback. Empty string if not found.  

## Example
Default
```
server.Use(realip.Middleware())
```
With Config
```
conf := realip.Config()
server.Use()
```
   
Get Result
```
// RouteHandlerFunc
func(w dhttp.ResponseWriter, r *dhttp.Request) {
  ip := realip.Get(r)
  // do your handling
}
```