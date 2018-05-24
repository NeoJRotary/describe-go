# Middleware Auth
DHTTP server middleware to do authentication
   
## Intro
set your authentication function on all request

## Parameters
- autoReject `bool`   
  server response 401 automatically or not.
- authFunc `func (w *dhttp.ResponseWriter, r *dhttp.Request) bool`
  your function which return `bool`. You can response by yourself and set false to autoReject.

## Method
- Valid `(r *dhttp.Request)` `bool`    
  boolean that this request is valid or not 