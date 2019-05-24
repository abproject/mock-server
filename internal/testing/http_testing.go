package testing

import (
	"bytes"
	"encoding/json"
	. "github.com/abproject/mock-server/internal/router"
	"io"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type HttpTestCase struct {
	Type            string
	Path            string
	Headers         map[string]string
	Body            string
	ExpectedStatus  int
	ExpectedBody    string
	ExpectedHeaders map[string]string
	realStatus  int
	realBody    string
	realHeaders map[string]string
}

func RunCases(name string, tests *[]HttpTestCase, t *testing.T) {
	for i, testCase := range *tests {
		doRequest(&testCase)
		(*tests)[i] = testCase
	}
	DisplayTestCase(name, *tests, t)
}

func doRequest(testCase *HttpTestCase) {
	var reader io.Reader = nil
	if testCase.Body != "" {
		reader = strings.NewReader(testCase.Body)
	}
	url := "/" + testCase.Path
	requestType := strings.ToUpper(testCase.Type)
	request := httptest.NewRequest(requestType, url, reader)
	for headerKey, headerValue := range testCase.Headers {
		request.Header.Set(headerKey, headerValue)
	}

	response := httptest.NewRecorder()
	Router(response, request)

	res := response.Result()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := buf.String()

	testCase.realStatus = res.StatusCode
	testCase.realBody = string(body)
	headers := make(map[string]string)
	for headerKey, headerValues := range res.Header {
		headers[headerKey] = strings.Join(headerValues[:],",")
	}
	testCase.realHeaders = headers
	if testCase.ExpectedHeaders == nil {
		testCase.ExpectedHeaders = make(map[string]string)
	}
}

func DisplayTestCase(name string, tests []HttpTestCase, t *testing.T) {
	for _, test := range tests {
		var realBody interface{}
		var expectedBody interface{}
		json.Unmarshal([]byte(test.realBody), &realBody)
		json.Unmarshal([]byte(test.ExpectedBody), &expectedBody)
		if test.realStatus != test.ExpectedStatus ||
			//test.realBody != test.ExpectedBody ||
			!reflect.DeepEqual(realBody, expectedBody) ||
			!reflect.DeepEqual(test.realHeaders, test.ExpectedHeaders) {
			t.Errorf(`
Failed HttpTestCase - %s:
Request: '%s'
Path:    '%s',
Headers: %#v,
Body:    '%s',
Response:
	Status %d expected to be %d
	Body '%s' expected to be '%s''
	Headers %#v expected to be %#v`,
				name,
				test.Type,
				test.Path,
				test.Headers,
				test.Body,
				test.realStatus,
				test.ExpectedStatus,
				test.realBody,
				test.ExpectedBody,
				test.realHeaders,
				test.ExpectedHeaders)
		}
	}
}
