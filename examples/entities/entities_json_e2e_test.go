package examplesentities

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/abproject/mock-server/internal/file"
	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/parser"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func TestEntitiesMockJSONE2E(t *testing.T) {
	router := configureJSON(t)
	testCases := GetEntitiesMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestEntitiesJSONE2E(t *testing.T) {
	router := configureJSON(t)
	testCases := GetEntitiesAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func configureJSON(t *testing.T) router.IRouter {
	logger := log.New(os.Stdout, "json e2e ", log.LstdFlags|log.Lshortfile)
	restStorage := rest.MakeStorage()
	fileStorage := file.MakeStorage()

	path, _ := filepath.Abs("../..")
	parserContext := models.AppContext{
		Logger:      logger,
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
		Path:        path,
	}
	parser := parser.New(parserContext)
	parser.Parse("config.json")

	routerContext := models.AppContext{
		Logger:      logger,
		RestStorage: &restStorage,
		FileStorage: &fileStorage,
	}
	router := router.New(routerContext)

	return router
}
