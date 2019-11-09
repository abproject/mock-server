package sharedhttphandler

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	sharedcomparator "github.com/abproject/mock-server/shared/comparator"
)

type testCase struct {
	comparator sharedcomparator.Comparator
	request    HTTPRequest
	response   HTTPResponse
	router     func(testCase testCase) func(w http.ResponseWriter, r *http.Request)
}

func TestHTTPHandler(t *testing.T) {
	testCases := []testCase{
		{
			comparator: sharedcomparator.NewTestingComparator(t, "GET request"),
			request: HTTPRequest{
				Method: "GET",
				URL:    "/URL",
				Headers: map[string]string{
					"header-get": "header-get-value",
				},
			},
			response: HTTPResponse{
				Status: 200,
				Body:   []byte("<html><head></head><body>OK</body></html>"),
				Headers: map[string]string{
					"Content-Type": "text/html; charset=utf-8",
				},
			},
			router: func(testCase testCase) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					compareRequest(testCase, r)

					w.Header().Add("Content-Type", "text/html")
					w.Header().Add("Content-Type", "charset=utf-8")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("<html><head></head><body>OK</body></html>"))
				}
			},
		},
		{
			comparator: sharedcomparator.NewTestingComparator(t, "POST request"),
			request: HTTPRequest{
				Method: "POST",
				URL:    "/URL",
				Body:   []byte("body"),
				Headers: map[string]string{
					"header-get": "header-get-value",
				},
			},
			response: HTTPResponse{
				Status: 201,
				Body:   nil,
				Headers: map[string]string{
					"Content-Type": "text/html; charset=utf-8",
				},
			},
			router: func(testCase testCase) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					compareRequest(testCase, r)

					w.Header().Add("Content-Type", "text/html; charset=utf-8")
					w.WriteHeader(http.StatusCreated)
				}
			},
		},
	}

	for _, testCase := range testCases {
		actualResponse := SendHTTPRequest(testCase.router(testCase), &testCase.request)
		compareResponse(testCase, actualResponse)
	}
}

func compareRequest(testCase testCase, r *http.Request) {
	compareRequestMethod(testCase, r)
	compareRequestHeaders(testCase, r)
	compareRequestBody(testCase, r)
}

func compareResponse(testCase testCase, actualResponse HTTPResponse) {
	compareResponseStatus(testCase, actualResponse)
	compareResponseHeaders(testCase, actualResponse)
	compareResponseBody(testCase, actualResponse)
}

func compareRequestMethod(testCase testCase, r *http.Request) {
	testCase.comparator.Equal(strings.ToUpper(testCase.request.Method), strings.ToUpper(r.Method), "Wrong Request Method")
}

func compareRequestHeaders(testCase testCase, r *http.Request) {
	normilizedHeaders := normilizeHeaders(r.Header)
	testCase.comparator.Equal(testCase.request.Headers, normilizedHeaders, "Wrong Request Headers")
}

func normilizeHeaders(headers map[string][]string) map[string]string {
	lowerHeaders := make(map[string]string)
	for key, values := range headers {
		normilizedKey := strings.ToLower(key)
		normilizedHeader := strings.Join(values[:], ";")
		lowerHeaders[normilizedKey] = normilizedHeader
	}
	return lowerHeaders
}

func compareRequestBody(testCase testCase, r *http.Request) {
	var expectedRequestBody = []byte("")
	if testCase.request.Body != nil {
		expectedRequestBody = testCase.request.Body
	}
	requestBody, _ := ioutil.ReadAll(r.Body)
	testCase.comparator.Equal(expectedRequestBody, requestBody, "Wrong Request Body")
}

func compareResponseStatus(testCase testCase, actualResponse HTTPResponse) {
	testCase.comparator.Equal(testCase.response.Status, actualResponse.Status, "Wrong Response Status")
}

func compareResponseHeaders(testCase testCase, actualResponse HTTPResponse) {
	testCase.comparator.Equal(testCase.response.Headers, actualResponse.Headers, "Wrong Response Headers")
}

func compareResponseBody(testCase testCase, actualResponse HTTPResponse) {
	var expectedResponseBody = []byte("")
	if testCase.response.Body != nil {
		expectedResponseBody = testCase.response.Body
	}
	testCase.comparator.Equal(expectedResponseBody, actualResponse.Body, "Wrong Response Body")
}
