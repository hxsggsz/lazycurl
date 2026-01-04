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
	req, err := http.NewRequest(ro.Method, ro.Url, nil)
	if err != nil {
		return Response{StatusCode: 500, Body: "Request Error: " + err.Error()}
	}

	req.Header.Set("Content-Type", "application/json")
	// log.Println("url: ", ro.Url, "req url: ", req.URL)

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return Response{StatusCode: 500, Body: fmt.Sprintf("Network Error: %s", err.Error())}
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	return Response{res.StatusCode, ro.Headers, string(body)}
}
