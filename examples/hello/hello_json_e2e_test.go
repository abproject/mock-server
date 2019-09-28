package exampleshello

import (
	"log"
	"os"
	"testing"

	"github.com/abproject/mock-server/internal/parser"
	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/internal/router"
)

func TestHelloMockJsonE2E(t *testing.T) {
	router := configureJSON(t)
	testCases := GetHelloMockCases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func TestHelloJsonE2E(t *testing.T) {
	router := configureJSON(t)
	testCases := GetHelloAPICases(t)

	for _, testCase := range testCases {
		response, request := testCase.TransformToHTTPResponseRequest()
		router.Route(response, request)
		testCase.AssertEquals(response)
	}
}

func configureJSON(t *testing.T) router.IRouter {
	logger := log.New(os.Stdout, "json e2e ", log.LstdFlags|log.Lshortfile)
	restStorage := rest.MakeStorage()

	parserContext := parser.Context{
		Logger:      logger,
		RestStorage: &restStorage,
	}
	parser := parser.New(parserContext)
	parser.Parse("config.json")

	routerContext := router.Context{
		Logger:      logger,
		RestStorage: &restStorage,
	}
	router := router.New(routerContext)

	return router
}
