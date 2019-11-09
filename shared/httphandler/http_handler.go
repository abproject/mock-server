package httphandler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func SendHttpRequest(router HttpRouter, httpRequest *HttpRequest) HttpResponse {
	request := prepareRequest(httpRequest)
	response := makeRequest(router, request)
	return parseResponse(response)
}

func prepareRequest(httpRequest *HttpRequest) *http.Request {
	reader := bytes.NewReader(httpRequest.Body)
	request := httptest.NewRequest(httpRequest.Method, httpRequest.URL, reader)
	for header, value := range httpRequest.Headers {
		request.Header.Add(header, value)
	}
	return request
}

func makeRequest(router HttpRouter, request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	router(response, request)
	return response
}

func parseResponse(response *httptest.ResponseRecorder) HttpResponse {
	result := response.Result()

	return HttpResponse{
		Status:  parseStatus(result),
		Body:    parseBody(result),
		Headers: parseHeaders(result),
	}
}

func parseStatus(result *http.Response) int {
	return result.StatusCode
}

func parseBody(result *http.Response) []byte {
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func parseHeaders(result *http.Response) map[string]string {
	headers := make(map[string]string)
	for key, header := range result.Header {
		headers[key] = strings.Join(header[:], "; ")
	}
	return headers
}
