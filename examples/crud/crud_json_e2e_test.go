package examplescrud

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/abproject/mock-server/internal/file"
	"github.com/abproject/mock-server/internal/parser"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func TestCRUDMockJsonE2E(t *testing.T) {
	router := configureJson(t)
	testCases := GetCRUDMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestCRUDJsonE2E(t *testing.T) {
	router := configureJson(t)
	testCases := GetCRUDAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func configureJson(t *testing.T) router.IRouter {
	logger := log.New(os.Stdout, "json e2e ", log.LstdFlags|log.Lshortfile)
	restStorage := rest.MakeStorage()
	fileStorage := file.MakeStorage()

	path, _ := filepath.Abs("../..")
	parserContext := parser.Context{
		Logger:      logger,
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
		Path:        path,
	}
	parser := parser.New(parserContext)
	parser.Parse("config.json")

	routerContext := router.Context{
		Logger:      logger,
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)

	return router
}
