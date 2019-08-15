package hello

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func configureAPI(t *testing.T) router.IRouter {
	restStorage := rest.MakeStorage()
	routerContext := router.Context{
		Logger:      log.New(os.Stdout, "api e2e ", log.LstdFlags|log.Lshortfile),
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)

	file, err := ioutil.ReadFile("config-api.json")
	if err != nil {
		t.Fatal(err)
	}
	reader := bytes.NewReader(file)
	request := httptest.NewRequest("POST", "/_api/rest", reader)
	response := httptest.NewRecorder()

	router.Route(response, request)
	return router
}

func TestHelloMockApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetHelloMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestHelloApiE2E(t *testing.T) {
	router := configureAPI(t)
	testCases := GetHelloAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}
