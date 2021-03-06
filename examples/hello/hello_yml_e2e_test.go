package exampleshello

import (
	"log"
	"os"
	"testing"

	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/parser"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func TestHelloMockYmlE2E(t *testing.T) {
	router := configureYml(t)
	testCases := GetHelloMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestHelloYmlE2E(t *testing.T) {
	router := configureYml(t)
	testCases := GetHelloAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func configureYml(t *testing.T) router.IRouter {
	logger := log.New(os.Stdout, "yaml e2e ", log.LstdFlags|log.Lshortfile)
	restStorage := rest.MakeStorage()

	parserContext := models.AppContext{
		Logger:      logger,
		RestStorage: &restStorage,
	}
	parser := parser.New(parserContext)
	parser.Parse("config.yaml")

	routerContext := models.AppContext{
		Logger:      logger,
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)

	return router
}
