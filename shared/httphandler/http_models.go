package sharedhttphandler

import "net/http"

type HTTPRouter func(w http.ResponseWriter, r *http.Request)

type HTTPRequest struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string]string
}

type HTTPResponse struct {
	Status  int
	Body    []byte
	Headers map[string]string
}
