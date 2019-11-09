package sharedhttphandler

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"strings"
)

func SendHTTPRequest(router HTTPRouter, httpRequest *HTTPRequest) HTTPResponse {
	reader := bytes.NewReader(httpRequest.Body)
	request := httptest.NewRequest(httpRequest.Method, httpRequest.URL, reader)
	for header, value := range httpRequest.Headers {
		request.Header.Add(header, value)
	}
	response := httptest.NewRecorder()
	router(response, request)
	r := response.Result()

	status := r.StatusCode

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	headers := make(map[string]string)
	for k, v := range r.Header {
		headers[k] = strings.Join(v[:], "; ")
	}

	return HTTPResponse{
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}
