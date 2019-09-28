package examplesapifiles

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/abproject/mock-server/internal/file"
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

func configureAPI(t *testing.T) router.IRouter {
	fileStorage := file.MakeStorage()
	routerContext := router.Context{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)
	return router
}

func configureAPIWithEntry(t *testing.T) (router.IRouter, string) {
	fileStorage := file.MakeStorage()
	routerContext := router.Context{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)

	fileName := "api-post-request.txt"
	path, _ := filepath.Abs(".")
	responseFile := test.SendFile(t, router, path, fileName)

	return router, responseFile.ID
}

func TestApiFilePostE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:                "POST Should get valid response",
				Method:              "POST",
				Status:              201,
				URI:                 "/_api/files",
				RequestFile:         "api-post-request.txt",
				RequestFileIsSource: true,
				ResponseFile:        "api-post-response.json",
				ExpectedResponse:    file.File{},
				ActualResponse:      file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileGetAllE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return empty array",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/files",
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []file.File{},
				ActualResponse:   []file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileGetAllWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL Should return one entry",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/files",
				RequestFile:      "",
				ResponseFile:     "api-get-all-response.json",
				ExpectedResponse: []file.File{},
				ActualResponse:   []file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileDeleteAllWithEntryE2E(t *testing.T) {
	router, _ := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE ALL Should delete all entries",
				Method:           "DELETE",
				Status:           204,
				URI:              "/_api/files",
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
				URI:              "/_api/files",
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []file.File{},
				ActualResponse:   []file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileGetByIdWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET by ID should return entry",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/files/" + id,
				RequestFile:      "",
				ResponseFile:     "api-get-response.json",
				ExpectedResponse: file.File{},
				ActualResponse:   file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileGetByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET by wrong ID should return error",
				Method:           "GET",
				Status:           404,
				URI:              "/_api/files/wrong-id",
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

func TestApiFilePutE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:                "PUT should create an entity",
				Method:              "PUT",
				Status:              200,
				URI:                 "/_api/files/some-id",
				RequestFile:         "api-put-request.txt",
				RequestFileIsSource: true,
				ResponseFile:        "api-put-response.json",
				ExpectedResponse:    file.File{},
				ActualResponse:      file.File{},
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL should return modified entries",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/files",
				RequestFile:      "",
				ResponseFile:     "api-get-all-modified-response.json",
				ExpectedResponse: []file.File{},
				ActualResponse:   []file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFilePutByWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:                "PUT by ID should modify entry",
				Method:              "PUT",
				Status:              200,
				URI:                 "/_api/files/" + id,
				RequestFile:         "api-put-request.txt",
				RequestFileIsSource: true,
				ResponseFile:        "api-put-response.json",
				ExpectedResponse:    file.File{},
				ActualResponse:      file.File{},
			}),
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET ALL should return modified entries",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/files",
				RequestFile:      "",
				ResponseFile:     "api-get-all-modified-response.json",
				ExpectedResponse: []file.File{},
				ActualResponse:   []file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileDeleteByWithEntryE2E(t *testing.T) {
	router, id := configureAPIWithEntry(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE by ID should delete entry",
				Method:           "DELETE",
				Status:           204,
				URI:              "/_api/files/" + id,
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
				URI:              "/_api/files",
				RequestFile:      "",
				ResponseFile:     "api-get-all-empty-response.json",
				ExpectedResponse: []file.File{},
				ActualResponse:   []file.File{},
			}),
	}

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestApiFileDeleteByWrongIdE2E(t *testing.T) {
	router := configureAPI(t)
	testCase := test.RestAPTestCaseFactory(t)
	testCases := []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "DELETE by wrong ID should return error",
				Method:           "DELETE",
				Status:           404,
				URI:              "/_api/files/wrong-id",
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
