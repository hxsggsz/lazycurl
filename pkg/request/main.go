package request

import (
	"io"
	"log"
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
	method  string
	url     string
	headers map[string]string
	body    string
}

type Response struct {
	StatusCode int
	Headers    map[string][]string
	Body       string
}

func RequestBuilder() *RequestOptions {
	return &RequestOptions{}
}

func (ro *RequestOptions) SetMethod(method string) *RequestOptions {
	ro.method = method
	return ro
}

func (ro *RequestOptions) SetURL(url string) *RequestOptions {
	ro.url = url
	return ro
}

func (ro *RequestOptions) Build() Response {
	req, _ := http.NewRequest(
		ro.method,
		ro.url,
		nil,
	)

	// TODO: set headers in view here
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)

	if err != nil {
		log.Println("Erro:", err)
		return Response{res.StatusCode, res.Header, err.Error()}
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	// how to get the headers in the response
	// res.Header

	return Response{res.StatusCode, res.Header, string(body)}
}
