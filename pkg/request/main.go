package request

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
)

type setHeader = func(key, value string)

type RequestOptions struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func RequestBuilder(method, url, body string, headers map[string]string) *RequestOptions {
	return &RequestOptions{Url: url, Method: method, Body: body, Headers: headers}
}

func (ro *RequestOptions) Send() Response {
	if ro.Url == "" {
		return Response{StatusCode: 500, Body: "Missing URL"}
	}

	req, err := http.NewRequest(ro.Method, ro.Url, nil)
	if err != nil {
		return Response{StatusCode: 500, Body: "Request Error: " + err.Error()}
	}

	setReqHeaders(ro.Headers, req.Header.Set)

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return Response{StatusCode: 500, Body: fmt.Sprintf("Network Error: %s", err.Error())}
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	simpleHeaders := simplifyHeaders(res.Header)

	return Response{res.StatusCode, simpleHeaders, string(body)}
}

func setReqHeaders(headers map[string]string, setHeader setHeader) {
	for k, v := range headers {
		if k != "" || v != "" {
			setHeader(k, v)
		}
	}
	setHeader("Content-Type", "application/json")
	setHeader("User-Agent", "LazyCurl/1.0 (Terminal API Client)")
}

func simplifyHeaders(headers http.Header) map[string]string {
	simpleHeaders := make(map[string]string)

	for name, values := range headers {
		headerValue := ""
		for i, v := range values {
			if i > 0 {
				headerValue += ", "
			}
			headerValue += v
		}
		simpleHeaders[name] = headerValue
	}
	return simpleHeaders
}
