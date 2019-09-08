package files

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
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

func TestFilesMockApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetFilesMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestFilesApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetFilesAPICases(t)

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

	configureFile(t, router, "examples/files/hello.txt", "examples/files/config-api-txt.json")
	configureFile(t, router, "examples/files/hello.json", "examples/files/config-api-json.json")

	return router
}

func configureFile(t *testing.T, router router.IRouter, dataFile string, configFile string) {
	path, _ := filepath.Abs("./../..")
	responseFile := test.SendFile(t, router, path, dataFile)

	file, err := ioutil.ReadFile(filepath.Join(path, configFile))
	if err != nil {
		t.Fatal(err)
	}
	data := rest.EndpointRestDto{}
	_ = json.Unmarshal([]byte(file), &data)
	data.Response.BodyFile = responseFile.ID

	requestByte, _ := json.Marshal(data)
	requestReader := bytes.NewReader(requestByte)

	request := httptest.NewRequest("POST", "/_api/rest/endpoint", requestReader)
	response := httptest.NewRecorder()
	router.Route(response, request)
}
