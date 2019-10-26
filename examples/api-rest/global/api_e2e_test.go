package examplesapirestglobal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/rest/restmodels"
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

var baseURL = "/_api/rest/global"

func configureAPI(t *testing.T) router.IRouter {
	restStorage := rest.MakeStorage()
	routerContext := router.Context{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)
	return router
}

func configureAPIWithEntry(t *testing.T) (router.IRouter, string) {
	restStorage := rest.MakeStorage()
	routerContext := router.Context{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)

	file, err := ioutil.ReadFile("api-post-request.json")
	if err != nil {
		t.Fatal(err)
	}
	reader := bytes.NewReader(file)
	request := httptest.NewRequest("POST", baseURL, reader)
	response := httptest.NewRecorder()
	router.Route(response, request)

	configureBody := restmodels.EndpointRestDto{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Result().Body)
	err = json.Unmarshal(buf.Bytes(), &configureBody)
	if err != nil {
		t.Fatal(err)
	}

	return router, configureBody.ID
}

func TestApiRestGlobalPostE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "POST Should get valid response",
				Method:           "POST",
				Status:           201,
				URI:              baseURL,
				RequestFile:      "api-post-request.json",
				ResponseFile:     "api-post-response.json",
				ExpectedResponse: restmodels.EndpointRestDto{},
				ActualResponse:   restmodels.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGlobalGetE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET Should return empty global configuration",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-empty-response.json",
				ExpectedResponse: restmodels.EndpointRestDto{},
				ActualResponse:   restmodels.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGlobalGetWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET Should return global configuration",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-response.json",
				ExpectedResponse: restmodels.EndpointRestDto{},
				ActualResponse:   restmodels.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGlobalDeleteE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE Should delete event if not exist",
				Method:           "DELETE",
				Status:           204,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "",
				ExpectedResponse: nil,
				ActualResponse:   nil,
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGlobalDeleteWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE Should delete global configuration",
				Method:           "DELETE",
				Status:           204,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "",
				ExpectedResponse: nil,
				ActualResponse:   nil,
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET Should return empty global configuration",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-empty-response.json",
				ExpectedResponse: restmodels.EndpointRestDto{},
				ActualResponse:   restmodels.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}
