# Middleware Access
DHTTP server middleware to access request by IP/CIDR    

## Intro
Response 403 Forbidden if request is not allowed. It check by order of Rules slice, stop when match.　

## Config
- UseRealIP `bool`   
  By default it check by `http.Request.RemoteAddr`. It can use RealIP if you `server.Use(realip.Middleware())` first.
- Rules `[]string`    
  access rules. `allow 1.2.3.4` or `deny 10.0.0.0/16`. You can use method to create it.

## Allow / Deny
String can be IP, CIDR, `all`, `loopback`    

## Method
- Allow `(s string)` `string`   
  get rule string of allowance.
- Deny `(s string)` `string`    
  get rule string of denial.
- `(*Config)` Allow `(s string)`   
  add allow rule in Config
- `(*Config)` Deny `(s string)`   
  add deny rule in Config

## Example
```
conf := access.Config{
  UseRealIP: true,
  Rules: []string{
    access.Allow("10.0.0.0/16"),
    access.Deny("all"),
  },
} 
```
or
```
conf := access.Config{}
conf.Allow("1.1.1.1")
conf.Deny("2.3.0.0/24")
conf.Allow("all")
```
Use 
```
server.Use(access.Middleware(conf))
```
