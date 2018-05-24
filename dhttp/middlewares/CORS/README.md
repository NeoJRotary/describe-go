# Middleware CORS
DHTTP server middleware for CORS   

## Intro
Add CORS headers. It response automatically if get `OPTIONS` method request.   

## Config
- NeedOrigin `bool`   
  default true. Only setup CORS header when request header has `Origin` when method is not OPTIONS.
- Origins `[]string`     
  default return `*`
- Methods `[]string`     
  default is `GET`, `POST`, `HEAD`
- Headers `[]string`     
  default return same as `Access-Control-Request-Headers`

## Exampl
Default
```
server.Use(cors.Middleware())
```
With Config
```
conf := cors.Config{
  NeedOrigin: true,
  Origins: []string{"mydomain.com", "sub.mydomain.com"},
  Methods: []string{"GET", "POST", "PUT", "DELETE"},
}

server.Use(cors.Middleware(conf))
```