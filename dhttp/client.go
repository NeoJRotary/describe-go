package dhttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	D "github.com/NeoJRotary/describe-go"
)

// TypeClient http function collections struct
type TypeClient struct {
	err     error
	Request *http.Request

	Method      string
	URL         string
	ContentType string
	Body        []byte
	BodyReader  io.Reader
	Header      map[string]string
	Query       map[string]string
	Timeout     time.Duration
}

// Response inheritance of http.Response
type Response struct {
	*http.Response
}

// Client get new Client, can be inited by TypeClient
func Client(input ...TypeClient) *TypeClient {
	c := TypeClient{
		Request: &http.Request{
			Method:     "GET",
			URL:        nil,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
			Body:       nil,
			Host:       "",
		},
	}

	if len(input) > 0 {
		init := input[0]
		c.Request.Method = init.Method
		c.SetURL(init.URL)
		c.SetContentType(init.ContentType)
		if init.Body != nil {
			c.SetBody(init.Body)
		}
		if init.BodyReader != nil {
			c.SetBodyReader(init.BodyReader)
		}

		for k, v := range init.Header {
			c.SetHeader(k, v)
		}

		for k, v := range init.Query {
			c.SetQuery(k, v)
		}
	}

	return &c
}

// GET set method to GET
func (c *TypeClient) GET() *TypeClient {
	c.Request.Method = http.MethodGet
	return c
}

// POST set method to POST
func (c *TypeClient) POST() *TypeClient {
	c.Request.Method = http.MethodPost
	return c
}

// PUT set method to PUT
func (c *TypeClient) PUT() *TypeClient {
	c.Request.Method = http.MethodPut
	return c
}

// PATCH set method to PATCH
func (c *TypeClient) PATCH() *TypeClient {
	c.Request.Method = http.MethodPatch
	return c
}

// DELETE set method to DELETE
func (c *TypeClient) DELETE() *TypeClient {
	c.Request.Method = http.MethodDelete
	return c
}

// OPTIONS set method to OPTIONS
func (c *TypeClient) OPTIONS() *TypeClient {
	c.Request.Method = http.MethodOptions
	return c
}

// CONNECT set method to CONNECT
func (c *TypeClient) CONNECT() *TypeClient {
	c.Request.Method = http.MethodConnect
	return c
}

// HEAD set method to HEAD
func (c *TypeClient) HEAD() *TypeClient {
	c.Request.Method = http.MethodHead
	return c
}

// TRACE set method to TRACE
func (c *TypeClient) TRACE() *TypeClient {
	c.Request.Method = http.MethodTrace
	return c
}

// SetContentType set content-type header
func (c *TypeClient) SetContentType(contentType string) *TypeClient {
	c.Request.Header.Set("Content-Type", contentType)
	return c
}

// SetURL set URL
func (c *TypeClient) SetURL(urlString string) *TypeClient {
	u, err := url.Parse(urlString)
	if D.IsErr(err) {
		c.err = err
		return c
	}
	c.Request.URL = u
	c.Request.Host = D.String(u.Host).TrimSuffix(":").Get()
	return c
}

// SetBody set http body by []byte
func (c *TypeClient) SetBody(buf []byte) *TypeClient {
	return c.SetBodyReader(bytes.NewBuffer(buf))
}

// SetBodyReader set http body by io.Reader
func (c *TypeClient) SetBodyReader(reader io.Reader) *TypeClient {
	rc, ok := reader.(io.ReadCloser)
	if !ok && reader != nil {
		rc = ioutil.NopCloser(reader)
	}
	c.Request.Body = rc

	if reader != nil {
		switch v := reader.(type) {
		case *bytes.Buffer:
			c.Request.ContentLength = int64(v.Len())
		case *bytes.Reader:
			c.Request.ContentLength = int64(v.Len())
		case *strings.Reader:
			c.Request.ContentLength = int64(v.Len())
		default:
		}
		if c.Request.ContentLength == 0 {
			c.Request.Body = http.NoBody
		}
	}

	return c
}

// SetHeader set one header key
func (c *TypeClient) SetHeader(key, value string) *TypeClient {
	c.Request.Header.Set(key, value)
	return c
}

// SetQuery set one query key
func (c *TypeClient) SetQuery(key, value string) *TypeClient {
	c.Request.URL.Query().Set(key, value)
	return c
}

// SetTimeout set timeout
func (c *TypeClient) SetTimeout(dur time.Duration) *TypeClient {
	c.Timeout = dur
	return c
}

// Do do http request
func (c *TypeClient) Do() (*Response, error) {
	if D.IsErr(c.err) {
		return nil, c.err
	}
	client := &http.Client{
		Timeout: c.Timeout,
	}

	res, err := client.Do(c.Request)
	if D.IsErr(err) {
		return nil, err
	}
	return &Response{res}, nil
}

// ReadAllBody read all response body
func (r *Response) ReadAllBody() []byte {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if D.IsErr(err) {
		return nil
	}
	return b
}
