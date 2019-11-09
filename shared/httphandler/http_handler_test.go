package httphandler

import (
	"net/http"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	testCaseFactory := httpHandlerTestCaseFactory(t)
	testCases := []httpHandlerTest{
		testCaseFactory(&httpHandlerTestCaseOptions{
			name: "GET request",
			request: HttpRequest{
				Method: "GET",
				URL:    "/url",
				Headers: map[string]string{
					"header-get": "header-get-value",
				},
			},
			response: HttpResponse{
				Status: 200,
				Body:   []byte("<html><head></head><body>OK</body></html>"),
				Headers: map[string]string{
					"Content-Type": "text/html; charset=utf-8",
				},
			},
			router: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "text/html")
				w.Header().Add("Content-Type", "charset=utf-8")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("<html><head></head><body>OK</body></html>"))

			},
		}),
		testCaseFactory(&httpHandlerTestCaseOptions{
			name: "POST request",
			request: HttpRequest{
				Method: "POST",
				URL:    "/url",
				Body:   []byte("body"),
				Headers: map[string]string{
					"header-get": "header-get-value",
				},
			},
			response: HttpResponse{
				Status: 201,
				Headers: map[string]string{
					"Content-Type": "text/html; charset=utf-8",
				},
			},
			router: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusCreated)
			},
		}),
	}

	for _, testCase := range testCases {
		actualResponse := testCase.makeRequest()
		testCase.compareResponse(actualResponse)
	}
}
