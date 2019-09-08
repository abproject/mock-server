package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

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
	request := httptest.NewRequest("POST", "/_api/rest/endpoint", reader)
	response := httptest.NewRecorder()
	router.Route(response, request)

	configureBody := rest.EndpointRestDto{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Result().Body)
	err = json.Unmarshal(buf.Bytes(), &configureBody)
	if err != nil {
		t.Fatal(err)
	}

	return router, configureBody.ID
}

func TestApiRestPostE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "POST Should get valid response",
				Method:           "POST",
				Status:           201,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "api-post-request.json",
				ResponseFile:     "api-post-response.json",
				ExpectedResponse: rest.EndpointRestDto{},
				ActualResponse:   rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGetAllE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return empty array",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGetAllWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return one entry",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "api-get-all-response.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestDeleteAllWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE ALL Should delete all entries",
				Method:           "DELETE",
				Status:           204,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "",
				ExpectedResponse: nil,
				ActualResponse:   nil,
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return empty array",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGetByIdWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET by ID should return entry",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint/" + id,
				RequestFile:      "",
				ResponseFile:     "api-get-response.json",
				ExpectedResponse: rest.EndpointRestDto{},
				ActualResponse:   rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestGetByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET by wrong ID should return error",
				Method:           "GET",
				Status:           404,
				URI:              "/_api/rest/endpoint/wrong-id",
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

func TestApiRestPutByWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "PUT by ID should modify entry",
				Method:           "PUT",
				Status:           200,
				URI:              "/_api/rest/endpoint/" + id,
				RequestFile:      "api-put-request.json",
				ResponseFile:     "api-put-response.json",
				ExpectedResponse: rest.EndpointRestDto{},
				ActualResponse:   rest.EndpointRestDto{},
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL should return modified entries",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "api-get-all-modified-response.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestPutByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "PUT by wrong ID should return error",
				Method:           "PUT",
				Status:           404,
				URI:              "/_api/rest/endpoint/wrong-id",
				RequestFile:      "api-put-request.json",
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

func TestApiRestDeleteByWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE by ID should delete entry",
				Method:           "DELETE",
				Status:           204,
				URI:              "/_api/rest/endpoint/" + id,
				RequestFile:      "",
				ResponseFile:     "",
				ExpectedResponse: nil,
				ActualResponse:   nil,
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return empty entries",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestDeleteByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE by wrong ID should return error",
				Method:           "DELETE",
				Status:           404,
				URI:              "/_api/rest/endpoint/wrong-id",
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
