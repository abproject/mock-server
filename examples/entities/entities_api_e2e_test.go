package examplesentities

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
	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
	"github.com/abproject/mock-server/test"
)

func TestEntitiesMockApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetEntitiesMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestEntitiesApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetEntitiesAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func configureAPI(t *testing.T) router.IRouter {
	restStorage := rest.MakeStorage()
	fileStorage := file.MakeStorage()
	routerContext := models.AppContext{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)

	fileAll := configureFile(t, router, "examples/entities/data-all.json")
	fileOne := configureFile(t, router, "examples/entities/data-new.json")

	path, _ := filepath.Abs("./../..")

	file, err := ioutil.ReadFile(filepath.Join(path, "examples/entities/config-api.json"))
	if err != nil {
		t.Fatal(err)
	}
	dto := models.EntityRestDto{}
	_ = json.Unmarshal([]byte(file), &dto)

	dto.Data = fileAll.ID
	dto.NewEntity = fileOne.ID

	requestByte, _ := json.Marshal(dto)
	requestReader := bytes.NewReader(requestByte)

	request := httptest.NewRequest("POST", "/_api/rest/entities", requestReader)
	response := httptest.NewRecorder()
	router.Route(response, request)

	return router
}

func configureFile(t *testing.T, router router.IRouter, dataFile string) models.File {
	path, _ := filepath.Abs("./../..")
	return test.SendFile(t, router, path, dataFile)
}
