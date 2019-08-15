package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

// RestAPITestCaseConfig HTTP Test Case Config
type RestAPITestCaseConfig struct {
	Name             string
	Method           string
	Status           int
	URI              string
	RequestFile      string
	ResponseFile     string
	ExpectedResponse interface{}
	ActualResponse   interface{}
}

// RestAPITestCase HTTP Test Case
type RestAPITestCase struct {
	RestAPITestCaseConfig
	t           *testing.T
	errorHolder ErrorHolder
}

// RestAPTestCaseFactory HTTP configuration: Request and Response
func RestAPTestCaseFactory(t *testing.T) func(config *RestAPITestCaseConfig) RestAPITestCase {
	return func(config *RestAPITestCaseConfig) RestAPITestCase {
		return RestAPITestCase{
			RestAPITestCaseConfig: *config,
			t:                     t,
			errorHolder:           NewErrorHolder(),
		}
	}
}

// TransformToHTTPResponseRequest Generate Request and Response for Router
func (testCase RestAPITestCase) TransformToHTTPResponseRequest() (*httptest.ResponseRecorder, *http.Request) {
	response := httptest.NewRecorder()
	if testCase.Method == "GET" || testCase.Method == "DELETE" {
		request := httptest.NewRequest(testCase.Method, testCase.URI, nil)
		return response, request
	}
	file, err := os.Open(testCase.RequestFile)
	if err != nil {
		requestFileGroupName := fmt.Sprintf("RequestFile: %s", testCase.RequestFile)
		requestFileGroup := testCase.errorHolder.Group(requestFileGroupName)
		errorMessage := fmt.Sprintf("File open error:\n%+v", err)
		requestFileGroup(errorMessage, 1)
	}
	request := httptest.NewRequest(testCase.Method, testCase.URI, file)
	return response, request
}

// AssertEquals checking Router Response
func (testCase *RestAPITestCase) AssertEquals(response *httptest.ResponseRecorder) {
	t := testCase.t
	expectedBody, actualBody := testCase.validate(response)
	if expectedBody != nil || actualBody != nil {
		testCase.compareBody(expectedBody, actualBody)
	}

	if testCase.errorHolder.HasErrors() {
		name := fmt.Sprintf("TEST CASE: %s", testCase.Name)
		outline := strings.Repeat("-", len(name))
		t.Logf("\n%s\n%s\n%s\n", outline, name, outline)
		testCase.errorHolder.Print(t)
	}
}

func (testCase *RestAPITestCase) validate(response *httptest.ResponseRecorder) (interface{}, interface{}) {
	expectedStatusGroupName := fmt.Sprintf("Status")
	expectedStatusGroup := testCase.errorHolder.Group(expectedStatusGroupName)
	if response.Code != testCase.Status {
		errorMessage := fmt.Sprintf("Expected: %d\n\tActual: %d", testCase.Status, response.Code)
		expectedStatusGroup(errorMessage, 1)
	}

	expectedBodyGroupName := fmt.Sprintf("Expected Body")
	expectedBodyGroup := testCase.errorHolder.Group(expectedBodyGroupName)

	expectedBody := testCase.ExpectedResponse

	if testCase.ResponseFile != "" {
		file, err := ioutil.ReadFile(testCase.ResponseFile)
		if err != nil {
			errorMessage := fmt.Sprintf("File open error: %s\n%+v", testCase.ResponseFile, err)
			expectedBodyGroup(errorMessage, 1)
		}

		if expectedBody != nil {
			err = json.Unmarshal([]byte(file), &expectedBody)
			if err != nil {
				errorMessage := fmt.Sprintf("Couldn't Unmarshal to EndpointRestDto\n%+v", err)
				expectedBodyGroup(errorMessage, 1)
			}
		}
	}

	actualBodyGroupName := fmt.Sprintf("Actual Body")
	actualBodyGroup := testCase.errorHolder.Group(actualBodyGroupName)

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Result().Body)
	actualBody := testCase.ActualResponse
	if actualBody != nil {
		err := json.Unmarshal(buf.Bytes(), &actualBody)
		if err != nil {
			errorMessage := fmt.Sprintf("Couldn't Unmarshal to EndpointRestDto\n%+v", err)
			actualBodyGroup(errorMessage, 1)
		}
	}

	return expectedBody, actualBody
}

func (testCase *RestAPITestCase) compareBody(expected interface{}, actual interface{}) {
	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)
	if (expectedValue.Kind() == reflect.Slice || expectedValue.Kind() == reflect.Array) &&
		(actualValue.Kind() == reflect.Slice || actualValue.Kind() == reflect.Array) {

		actualSlice := make([]interface{}, actualValue.Len())
		for i := 0; i < actualValue.Len(); i++ {
			actualSlice[i] = actualValue.Index(i).Interface()
		}

		expectedSlice := make([]interface{}, expectedValue.Len())
		for i := 0; i < expectedValue.Len(); i++ {
			actualValueReflection := actualValue.Index(i).Interface()
			id := actualValueReflection.(map[string]interface{})["id"]
			value := expectedValue.Index(i).Interface()
			m, _ := value.(map[string]interface{})
			m["id"] = id
			expectedSlice[i] = m
		}

		compareBodyGroupName := fmt.Sprintf(`
		Expected Body: %+v
		Actual Body:   %+v`, expectedSlice, actualSlice)
		compareBodyGroup := testCase.errorHolder.Group(compareBodyGroupName)

		if !reflect.DeepEqual(expectedSlice, actualSlice) {
			compareBodyGroup("Not Equal Configurations", 1)
		}
	} else {
		id := actualValue.Interface().(map[string]interface{})["id"]
		value := expectedValue.Interface()
		expectedParsedValue, _ := value.(map[string]interface{})
		expectedParsedValue["id"] = id

		compareBodyGroupName := fmt.Sprintf(`
		Expected Body: %+v
		Actual Body:   %+v`, expectedParsedValue, actual)
		compareBodyGroup := testCase.errorHolder.Group(compareBodyGroupName)

		if !reflect.DeepEqual(expectedParsedValue, actual) {
			compareBodyGroup("Must be equal", 1)
		}
	}
}
