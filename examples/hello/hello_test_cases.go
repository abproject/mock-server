package hello

import (
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/test"
)

// GetHelloMockCases Returns all Test Cases
func GetHelloMockCases(t *testing.T) []test.RestMockTestCase {
	testCase := test.RestMockTestCaseFactory(t)

	return []test.RestMockTestCase{
		testCase(
			"GET should return correct body, headers ans status",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/hello",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "text/html",
				},
				Body: "Hello, World!",
			}),
		testCase(
			"GET should return NOT_FOUND error when given unknown path",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/unknown-path",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
			}),
		testCase(
			"POST should return NOT_FOUND error when only GET controller provided",
			test.RestMockTestCaseRequest{
				Type:    "POST",
				Path:    "/hello",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
			}),
		testCase(
			"PUT should return NOT_FOUND error when only GET controller provided",
			test.RestMockTestCaseRequest{
				Type:    "PUT",
				Path:    "/hello",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
			}),
		testCase(
			"DELETE should return NOT_FOUND error when only GET controller provided",
			test.RestMockTestCaseRequest{
				Type:    "DELETE",
				Path:    "/hello",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
			}),
	}
}

// GetHelloAPICases Returns all Test Cases
func GetHelloAPICases(t *testing.T) []test.RestAPITestCase {
	testCase := test.RestAPTestCaseFactory(t)

	return []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET should return all configuration",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoint",
				RequestFile:      "",
				ResponseFile:     "api-get-all.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}
}
