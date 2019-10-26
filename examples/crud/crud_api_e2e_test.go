package examplescrud

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/abproject/mock-server/internal/file"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/rest/restmodels"
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

func TestCRUDMockApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetCRUDMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestCRUDApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetCRUDAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func configureAPI(t *testing.T) router.IRouter {
	restStorage := rest.MakeStorage()
	fileStorage := file.MakeStorage()
	routerContext := router.Context{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)

	fileAll := configureFile(t, router, "examples/crud/data.json")
	fileOne := configureFile(t, router, "examples/crud/data-id-3.json")

	configureGlobal(t, router, "examples/crud/config-api-global.json")
	configureEndpoint(t, router, "examples/crud/config-api-get-all.json", fileAll.ID)
	configureEndpoint(t, router, "examples/crud/config-api-get.json", fileOne.ID)
	configureEndpoint(t, router, "examples/crud/config-api-post.json", fileOne.ID)
	configureEndpoint(t, router, "examples/crud/config-api-put.json", fileOne.ID)
	configureEndpoint(t, router, "examples/crud/config-api-delete.json", "")

	return router
}

func configureFile(t *testing.T, router router.IRouter, dataFile string) file.File {
	path, _ := filepath.Abs("./../..")
	return test.SendFile(t, router, path, dataFile)
}

func configureRequest(t *testing.T, router router.IRouter, configFile string, fileID string, url string) {
	path, _ := filepath.Abs("./../..")

	file, err := ioutil.ReadFile(filepath.Join(path, configFile))
	if err != nil {
		t.Fatal(err)
	}
	data := restmodels.EndpointRestDto{}
	_ = json.Unmarshal([]byte(file), &data)

	if fileID != "" {
		data.Response.BodyFile = fileID
	}

	requestByte, _ := json.Marshal(data)
	requestReader := bytes.NewReader(requestByte)

	request := httptest.NewRequest("POST", url, requestReader)
	response := httptest.NewRecorder()
	router.Route(response, request)
}

func configureEndpoint(t *testing.T, router router.IRouter, configFile string, fileID string) {
	configureRequest(t, router, configFile, fileID, "/_api/rest/endpoints")
}

func configureGlobal(t *testing.T, router router.IRouter, configFile string) {
	configureRequest(t, router, configFile, "", "/_api/rest/global")
}
