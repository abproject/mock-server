package httphandler

import "net/http"

type HttpRouter func(w http.ResponseWriter, r *http.Request)

type HttpRequest struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string]string
}

type HttpResponse struct {
	Status  int
	Body    []byte
	Headers map[string]string
}
