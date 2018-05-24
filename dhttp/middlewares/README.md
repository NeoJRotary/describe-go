# Middlewares
Do something before start handling       

## Intro
It can be set at `Server` as global middleware, or on `Route` as path middleware. Middleware can use other middleware's output if it is after target middleware. Execution order be decided by who be `Use()` first. All middlewares will be executed before your route handler.   
   
Middleware may response for you (like CORS, Auth). It is suggested that use middleware's function to read output. Any output will bind to `dhttp.Request.MiddlewareValues`, it's a `map[string]interface{}`.
   
We prepared some middlewares that maybe useful : 
- [Access](https://github.com/NeoJRotary/describe-go/blob/master/dhttp/middlewares/Access) : access request by IP/CIDR
- [Auth](https://github.com/NeoJRotary/describe-go/blob/master/dhttp/middlewares/Auth) : authentication request header/token/whatever
- [CORS](https://github.com/NeoJRotary/describe-go/blob/master/dhttp/middlewares/CORS) : add (or response) CORS header 
- [RealIP](https://github.com/NeoJRotary/describe-go/blob/master/dhttp/middlewares/RealIP) : get client real IP from header

## Method
- Rename `(string)`   
  rename this middleware to avoid value overwriting.

## How To Make
### Structure
In `Use()` function it required a pointer of `Middleware` type.   
```
type Middleware struct {
	Name    string
	Config  interface{}
	Handler MiddlewareFunc
}

type MiddlewareFunc func(w *ResponseWriter, r *Request, config interface{}) interface{}
```
- Name : name in values map
- Config : config pass to handler
- Handler : function return any value or do reponsing. Request pipeline will stop after middleware `Write` something.
   
### Methods
Preare 2 methods in your middleware pakage to initialize middleware.      
```
func Middleware(~some config~) *dhttp.Middleware {
  // init your config

  return &dhttp.Middleware{
    Name: "MyMiddleware",
    Config: conf,
    Handler: handler,
  }
}

// handler bind to object
func handler(w *dhttp.ResponseWriter, r *dhttp.Request, config interface{}) interface{}
```
Extra methods to let user get value easily. Plus more if you think user need!
```
func Get(r *dhttp.Request) string {
  return GetBy("MyMiddleware", r)
}

func GetBy(name string, r *dhttp.Request) string {
  v, ok := r.MiddlewareValues[name]
  if !ok {
    return ""
  }
  return v.(string)
}
```