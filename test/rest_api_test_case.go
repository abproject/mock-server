package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/abproject/mock-server/internal/rest"
)

// RestAPITestCase HTTP Test Case
type RestAPITestCase struct {
	name        string
	file        string
	path        string
	t           *testing.T
	errorHolder ErrorHolder
}

// RestAPTestCaseFactory HTTP configuration: Request and Response
func RestAPTestCaseFactory(t *testing.T) func(name string, file string, path string) RestAPITestCase {
	return func(name string, file string, path string) RestAPITestCase {
		return RestAPITestCase{
			name:        name,
			file:        file,
			path:        path,
			t:           t,
			errorHolder: NewErrorHolder(),
		}
	}
}

// TransformToHTTPResponseRequest Generate Request and Response for Router
func (testCase RestAPITestCase) TransformToHTTPResponseRequest() (*httptest.ResponseRecorder, *http.Request) {
	request := httptest.NewRequest("GET", testCase.path, nil)
	response := httptest.NewRecorder()
	return response, request
}

// AssertEquals checking Router Response
func (testCase *RestAPITestCase) AssertEquals(response *httptest.ResponseRecorder) {
	t := testCase.t
	expectedBody, actualBody := testCase.validate(response)
	testCase.compare(expectedBody, actualBody)

	if testCase.errorHolder.HasErrors() {
		name := fmt.Sprintf("TEST CASE: %s", testCase.name)
		outline := strings.Repeat("-", len(name))
		t.Logf("\n%s\n%s\n%s\n", outline, name, outline)
		testCase.errorHolder.Print(t)
	}
}

func (testCase *RestAPITestCase) validate(response *httptest.ResponseRecorder) ([]rest.EndpointRestDto, []rest.EndpointRestDto) {
	expectedBodyGroupName := fmt.Sprintf("Expected Body")
	expectedBodyGroup := testCase.errorHolder.Group(expectedBodyGroupName)

	file, err := ioutil.ReadFile(testCase.file)
	if err != nil {
		errorMessage := fmt.Sprintf("File open error: %s\n%+v", testCase.file, err)
		expectedBodyGroup(errorMessage, 1)
	}
	expectedBody := []rest.EndpointRestDto{}
	err = json.Unmarshal([]byte(file), &expectedBody)
	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't Unmarshal to EndpointRestDto\n%+v", err)
		expectedBodyGroup(errorMessage, 1)
	}

	actualBodyGroupName := fmt.Sprintf("Actual Body")
	actualBodyGroup := testCase.errorHolder.Group(actualBodyGroupName)

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Result().Body)
	actualBody := []rest.EndpointRestDto{}
	err = json.Unmarshal(buf.Bytes(), &actualBody)
	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't Unmarshal to EndpointRestDto\n%+v", err)
		actualBodyGroup(errorMessage, 1)
	}

	return expectedBody, actualBody
}

func (testCase *RestAPITestCase) compare(expected []rest.EndpointRestDto, actual []rest.EndpointRestDto) {
	compareBodyGroupName := fmt.Sprintf(`
	Expected Body: %+v
	Actual Body:   %+v`, expected, actual)
	compareBodyGroup := testCase.errorHolder.Group(compareBodyGroupName)

	if len(expected) != len(actual) {
		errorMessage := fmt.Sprintf("Not equal amount of Configurations\n\t\tExpected amount: %d\n\t\tActual amount: %d", len(expected), len(actual))
		compareBodyGroup(errorMessage, 1)
	} else {
		for i := 0; i < len(expected); i++ {
			expected[i].ID = actual[i].ID
		}
		if !reflect.DeepEqual(expected, actual) {
			compareBodyGroup("Not Equal Configurations", 1)
		}
	}
}
