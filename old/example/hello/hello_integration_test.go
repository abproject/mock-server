package hello

import (
	"bufio"
	"bytes"
	"github.com/abproject/mock-server/internal_/rest"
	"github.com/abproject/mock-server/internal_/router"
	httpTesting "github.com/abproject/mock-server/internal_/testing"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRequestInit(t *testing.T) {
	tests := []httpTesting.HttpTestCase{
		{
			Type: "GET",
			Path: "hello",
			ExpectedStatus: 200,
			ExpectedBody: "Hello, World!",
		},
		{
			Type: "GET",
			Path: "hello/world",
			ExpectedStatus: 404,
		},
		{
			Type: "POST",
			Path: "hello",
			ExpectedStatus: 404,
			ExpectedBody: "",
		},
		{
			Type: "GET",
			Path: "not-found",
			ExpectedStatus: 404,
			ExpectedBody: "",
		},
		{
			Type: "POST",
			Path: "not-found",
			ExpectedStatus: 404,
			ExpectedBody: "",
		},
	}
	//rest.Clear()
	//rest.FileRest(ParseConfig("config.yml").Rest)
	//httpTesting.RunCases("Hello (yml)", &tests, t)
	//
	//rest.Clear()
	//rest.FileRest(ParseConfig("config.json").Rest)
	//httpTesting.RunCases("Hello (json)", &tests, t)

	rest.Clear()
	file, _ := os.Open("/api.json")
	reader := bufio.NewReader(file)
	request := httptest.NewRequest("POST", "/_api/rest", reader)
	response := httptest.NewRecorder()
	router.Router(response, request)

	res := response.Result()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := buf.String()
	log.Printf(body)
	httpTesting.RunCases("Hello (API)", &tests, t)
}

