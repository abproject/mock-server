package examplesfiles

import (
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/test"
)

type body struct {
	Message string `json:"message" yaml:"message"`
}

// GetFilesMockCases Returns all Test Cases
func GetFilesMockCases(t *testing.T) []test.RestMockTestCase {
	testCase := test.RestMockTestCaseFactory(t)

	return []test.RestMockTestCase{
		testCase(
			"GET txt should return correct body, headers ans status",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/hello-txt",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "text/plain",
				},
				Body: "Hello from file!",
			}),
		testCase(
			"GET json should return correct body, headers ans status",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/hello-json",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: "{\"message\": \"Hello, World!\"}",
			}),
	}
}

// GetFilesAPICases Returns all Test Cases
func GetFilesAPICases(t *testing.T) []test.RestAPITestCase {
	testCase := test.RestAPTestCaseFactory(t)

	return []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET should return all configuration",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/endpoints",
				RequestFile:      "",
				ResponseFile:     "api-get-all.json",
				ExpectedResponse: []rest.EndpointRestDto{},
				ActualResponse:   []rest.EndpointRestDto{},
			}),
	}
}
