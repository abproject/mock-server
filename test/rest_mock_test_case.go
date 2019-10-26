package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

type void struct{}

var v void

var requestTypes = map[string]void{
	"GET":    v,
	"POST":   v,
	"PUT":    v,
	"DELETE": v,
}
var requestTypeKeys = getKeys(requestTypes)

func getKeys(map[string]void) string {
	var keys = make([]string, 0)
	for key := range requestTypes {
		keys = append(keys, key)
	}
	return strings.Join(keys, ", ")
}

// RestMockTestCaseRequest HTTP Request structure
type RestMockTestCaseRequest struct {
	Type    string
	Path    string
	Headers map[string]string
	Body    string
}

// RestMockTestCaseResponse HTTP Response structure
type RestMockTestCaseResponse struct {
	Status   int
	Headers  map[string]string
	Body     string
	BodyFile string
}

// RestMockTestCase HTTP Test Case
type RestMockTestCase struct {
	name        string
	request     RestMockTestCaseRequest
	response    RestMockTestCaseResponse
	t           *testing.T
	errorHolder ErrorHolder
}

// RestMockTestCaseFactory HTTP configuration: Request and Response
func RestMockTestCaseFactory(t *testing.T) func(name string, request RestMockTestCaseRequest, response RestMockTestCaseResponse) RestMockTestCase {
	return func(name string, request RestMockTestCaseRequest, response RestMockTestCaseResponse) RestMockTestCase {
		return RestMockTestCase{
			name:        name,
			request:     request,
			response:    response,
			t:           t,
			errorHolder: NewErrorHolder(),
		}
	}
}

// TransformToHTTPResponseRequest Generate Request and Response for Router
func (testCase RestMockTestCase) TransformToHTTPResponseRequest() (*httptest.ResponseRecorder, *http.Request) {
	r := testCase.request
	reader := strings.NewReader(testCase.request.Body)
	request := httptest.NewRequest(r.Type, r.Path, reader)
	for header, value := range r.Headers {
		request.Header.Set(header, value)
	}
	response := httptest.NewRecorder()
	return response, request
}

// AssertEquals checking Router Response
func (testCase *RestMockTestCase) AssertEquals(response *httptest.ResponseRecorder) {
	expected := testCase.response
	actual := response.Result()
	t := testCase.t

	testCase.validate()

	actualResponseGroupName := fmt.Sprintf(`
Expected Response: %+v
Actual Response:   %+v`, expected, *actual)
	actualResponseGroup := testCase.errorHolder.Group(actualResponseGroupName)

	if expected.Status != actual.StatusCode {
		errorMessage := fmt.Sprintf("Status Code:\n\t\tExpected: %d\n\t\tActual:   %d", expected.Status, actual.StatusCode)
		actualResponseGroup(errorMessage, 1)
	}

	actualHeader := make(map[string]string)
	for k, v := range actual.Header {
		actualHeader[k] = strings.Join(v[:], "; ")
	}
	if !reflect.DeepEqual(expected.Headers, actualHeader) {
		errorMessage := fmt.Sprintf("Headers:\n\t\tExpected: %+v\n\t\tActual:   %+v", expected.Headers, actualHeader)
		actualResponseGroup(errorMessage, 1)
	}

	var expectedBody string
	if expected.BodyFile != "" {
		file, err := ioutil.ReadFile(expected.BodyFile)
		if err != nil {
			errorMessage := fmt.Sprintf("File open error: %s\n%+v", expected.BodyFile, err)
			actualResponseGroup(errorMessage, 1)
		}
		expectedBody = string(file)
	} else {
		expectedBody = expected.Body
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(actual.Body)
	actualBody := buf.String()

	r, _ := regexp.Compile(`[^\s"]+|"([^"]*)"`)
	actualParts := r.FindAllString(actualBody, -1)
	expectedParts := r.FindAllString(expectedBody, -1)
	actualTrimmedBody := strings.Join(actualParts[:], "")
	expectedTrimmedBody := strings.Join(expectedParts[:], "")

	if expectedTrimmedBody != actualTrimmedBody {
		errorMessage := fmt.Sprintf("Body:\n\t\tExpected: %s\n\t\tActual:   %s", expectedTrimmedBody, actualTrimmedBody)
		actualResponseGroup(errorMessage, 1)
	}

	if testCase.errorHolder.HasErrors() {
		name := fmt.Sprintf("TEST CASE: %s", testCase.name)
		outline := strings.Repeat("-", len(name))
		t.Logf("\n%s\n%s\n%s\n", outline, name, outline)
		testCase.errorHolder.Print(t)
	}
}

func (testCase *RestMockTestCase) validate() {
	request := testCase.request
	response := testCase.response

	requestGroupName := fmt.Sprintf("Request: %+v", request)
	requestGroup := testCase.errorHolder.Group(requestGroupName)
	if _, exist := requestTypes[request.Type]; !exist {
		errorMessage := fmt.Sprintf("Type must be in range %s, but was: '%s'", requestTypeKeys, request.Type)
		requestGroup(errorMessage, 1)
	}
	if request.Path == "" {
		requestGroup("Path must not be empty", 1)
	}
	if request.Headers == nil {
		requestGroup("Headers must be initialised", 1)
	}

	responseGroupName := fmt.Sprintf("Response: %+v", response)
	responseGroup := testCase.errorHolder.Group(responseGroupName)
	if response.Status <= 0 {
		errorMessage := fmt.Sprintf("Status code must have positive value, was: %d", response.Status)
		responseGroup(errorMessage, 1)
	}
	if response.Headers == nil {
		responseGroup("Headers must be initialised", 1)
	}
}
