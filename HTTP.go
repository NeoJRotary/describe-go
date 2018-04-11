package describe

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// TypeHTTP http function collections struct
type TypeHTTP struct {
	Method     string
	Header     TypeHTTPKVMap
	Body       []byte
	URL        string
	Query      TypeHTTPKVMap
	StatusCode int
	RoutePath  string
	Request    *http.Request
	Response   *http.Response
	W          http.ResponseWriter
}

// TypeHTTPKVMap TypeHTTP Key-Value Map
type TypeHTTPKVMap map[string]string

// Get get value by key
func (m TypeHTTPKVMap) Get(key string) string {
	v, ok := m[key]
	if !ok {
		return ""
	}
	return v
}

// Set set value by key
func (m TypeHTTPKVMap) Set(key, value string) {
	m[key] = value
}

func converRequestToTypeHTTP(r *http.Request) *TypeHTTP {
	defer r.Body.Close()
	h := TypeHTTP{
		Method:  r.Method,
		Header:  map[string]string{},
		Query:   map[string]string{},
		Request: r,
	}
	b, err := ioutil.ReadAll(r.Body)
	if !IsErr(err) {
		h.Body = b
	}
	for k := range r.Header {
		h.Header[k] = r.Header.Get(k)
	}
	qv := r.URL.Query()
	for k := range qv {
		h.Query[k] = qv.Get(k)
	}
	return &h
}

// HTTP get *TypeHTTP
func HTTP() *TypeHTTP {
	return &TypeHTTP{
		Header: map[string]string{},
		Query:  map[string]string{},
	}
}

// Copy get copy of *TypeHTTP
func (h *TypeHTTP) Copy() *TypeHTTP {
	cp := &TypeHTTP{
		URL: h.URL,
	}
	for k, v := range h.Header {
		cp.Header[k] = v
	}
	for k, v := range h.Query {
		cp.Query[k] = v
	}
	copy(cp.Body, h.Body)
	return cp
}

// GET setup method GET
func (h *TypeHTTP) GET() *TypeHTTP {
	h.Method = "GET"
	return h
}

// POST setup method  POST
func (h *TypeHTTP) POST() *TypeHTTP {
	h.Method = "POST"
	return h
}

// PUT setup method PUT
func (h *TypeHTTP) PUT() *TypeHTTP {
	h.Method = "PUT"
	return h
}

// DELETE setup method DELETE
func (h *TypeHTTP) DELETE() *TypeHTTP {
	h.Method = "DELETE"
	return h
}

// PATCH setup method PATCH
func (h *TypeHTTP) PATCH() *TypeHTTP {
	h.Method = "PATCH"
	return h
}

// OPTIONS setup method OPTIONS
func (h *TypeHTTP) OPTIONS() *TypeHTTP {
	h.Method = "OPTIONS"
	return h
}

// HEAD setup method HEAD
func (h *TypeHTTP) HEAD() *TypeHTTP {
	h.Method = "HEAD"
	return h
}

// SetHeader set header
func (h *TypeHTTP) SetHeader(key, value string) *TypeHTTP {
	h.Header.Set(key, value)
	return h
}

// WithHeader replace header map
func (h *TypeHTTP) WithHeader(header map[string]string) *TypeHTTP {
	h.Header = header
	return h
}

// GetHeader get header by key
func (h *TypeHTTP) GetHeader(key string) string {
	return h.Header.Get(key)
}

// SetQuery set query
func (h *TypeHTTP) SetQuery(key, value string) *TypeHTTP {
	h.Query.Set(key, value)
	return h
}

// WithQuery replace header map
func (h *TypeHTTP) WithQuery(header map[string]string) *TypeHTTP {
	h.Header = header
	return h
}

// GetQuery get query by key
func (h *TypeHTTP) GetQuery(key string) string {
	return h.Query.Get(key)
}

// WithBody setup req body
func (h *TypeHTTP) WithBody(body []byte) *TypeHTTP {
	h.Body = body
	return h
}

// GetBody get req/res body
func (h *TypeHTTP) GetBody() []byte {
	return h.Body
}

// AtURL setup req url
func (h *TypeHTTP) AtURL(url string) *TypeHTTP {
	if !String(url).HasPrefix("http") {
		url = "http://" + url
	}
	h.URL = url
	return h
}

// Do do http request
func (h *TypeHTTP) Do() (typeH *TypeHTTP, err error) {
	ss := StringSlice(nil)
	for k, v := range h.Query {
		ss.Push(k + "=" + v)
	}
	// q := url.QueryEscape(ss.Join("&").Get())
	q := "?" + ss.Join("&").Get()
	client := &http.Client{}
	var req *http.Request
	if len(h.Body) == 0 {
		req, err = http.NewRequest(h.Method, h.URL+q, nil)
	} else {
		req, err = http.NewRequest(h.Method, h.URL+q, bytes.NewBuffer(h.Body))
	}
	if IsErr(err) {
		return nil, err
	}
	for k, v := range h.Header {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if IsErr(err) {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if IsErr(err) {
		return nil, err
	}

	typeH = &TypeHTTP{
		Body:       b,
		StatusCode: res.StatusCode,
		Response:   res,
	}

	return typeH, nil
}

// SetStatus set status code
func (h *TypeHTTP) SetStatus(code int) *TypeHTTP {
	h.StatusCode = code
	return h
}

// RelativePath get relative path from route path without prefix "/"
func (h *TypeHTTP) RelativePath() string {
	base := String(h.RoutePath).TrimPrefix("/").Get()
	return String(h.Request.URL.Path).TrimPrefix("/").TrimPrefix(base).TrimPrefix("/").Get()
}
