package examplesapirestentity

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

var baseURL = "/_api/rest/entities"

func configureAPI(t *testing.T) router.IRouter {
	restStorage := rest.MakeStorage()
	routerContext := models.AppContext{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)
	return router
}

func configureAPIWithEntry(t *testing.T) (router.IRouter, string) {
	restStorage := rest.MakeStorage()
	routerContext := models.AppContext{
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

	configureBody := models.EntityRestDto{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Result().Body)
	err = json.Unmarshal(buf.Bytes(), &configureBody)
	if err != nil {
		t.Fatal(err)
	}

	return router, configureBody.Name
}

func TestApiRestEndpointPostE2E(t *testing.T) {
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
				ExpectedResponse: models.EntityRestDto{},
				ActualResponse:   models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointGetAllE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return empty array",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []models.EntityRestDto{},
				ActualResponse:   []models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointGetAllWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return one entry",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-all-response.json",
				ExpectedResponse: []models.EntityRestDto{},
				ActualResponse:   []models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointDeleteAllWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE ALL Should delete all entries",
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
				Name:             "GET ALL Should return empty array",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []models.EntityRestDto{},
				ActualResponse:   []models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointGetByIdWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET by ID should return entry",
				Method:           "GET",
				Status:           200,
				URI:              baseURL + "/" + id,
				RequestFile:      "",
				ResponseFile:     "api-get-response.json",
				ExpectedResponse: models.EntityRestDto{},
				ActualResponse:   models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointGetByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET by wrong ID should return error",
				Method:           "GET",
				Status:           404,
				URI:              baseURL + "/wrong-id",
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

func TestApiRestEndpointPutByWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "PUT by ID should modify entry",
				Method:           "PUT",
				Status:           200,
				URI:              baseURL + "/" + id,
				RequestFile:      "api-put-request.json",
				ResponseFile:     "api-put-response.json",
				ExpectedResponse: models.EntityRestDto{},
				ActualResponse:   models.EntityRestDto{},
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL should return modified entries",
				Method:           "GET",
				Status:           200,
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-all-modified-response.json",
				ExpectedResponse: []models.EntityRestDto{},
				ActualResponse:   []models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointPutByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "PUT by wrong ID should return error",
				Method:           "PUT",
				Status:           404,
				URI:              baseURL + "/wrong-id",
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

func TestApiRestEndpointDeleteByWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE by ID should delete entry",
				Method:           "DELETE",
				Status:           204,
				URI:              baseURL + "/" + id,
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
				URI:              baseURL,
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []models.EntityRestDto{},
				ActualResponse:   []models.EntityRestDto{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiRestEndpointDeleteByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE by wrong ID should return error",
				Method:           "DELETE",
				Status:           404,
				URI:              baseURL + "/wrong-id",
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
