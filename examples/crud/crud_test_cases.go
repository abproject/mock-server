package examplescrud

import (
	"testing"

	"github.com/abproject/mock-server/internal/rest"
	"github.com/abproject/mock-server/test"
)

type body struct {
	ID         string   `json:"id" yaml:"id"`
	Name       string   `json:"name" yaml:"name"`
	Type       string   `json:"type" yaml:"type"`
	Period     int      `json:"period" yaml:"period"`
	Atmosphere []string `json:"atmosphere" yaml:"atmosphere"`
}

// GetCRUDMockCases Returns all Test Cases
func GetCRUDMockCases(t *testing.T) []test.RestMockTestCase {
	testCase := test.RestMockTestCaseFactory(t)

	return []test.RestMockTestCase{
		testCase(
			"GET ALL should return correct body, headers ans status",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/planets",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "data.json",
			}),
		testCase(
			"GET should return correct body, headers ans status when id=3",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/planets/3",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "data-id-3.json",
			}),
		testCase(
			"GET ALL should return correct body, headers ans status when id=1",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/planets/1",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "data-id-3.json",
			}),
		testCase(
			"POST should return correct body, headers ans status",
			test.RestMockTestCaseRequest{
				Type:    "POST",
				Path:    "/planets",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 201,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "data-id-3.json",
			}),
		testCase(
			"PUT should return correct body, headers ans status when id=3",
			test.RestMockTestCaseRequest{
				Type:    "PUT",
				Path:    "/planets/3",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "data-id-3.json",
			}),
		testCase(
			"PUT should return correct body, headers ans status when id=1",
			test.RestMockTestCaseRequest{
				Type:    "PUT",
				Path:    "/planets/1",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "data-id-3.json",
			}),
		testCase(
			"DELETE should return correct body, headers ans status when id=3",
			test.RestMockTestCaseRequest{
				Type:    "DELETE",
				Path:    "/planets/3",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 204,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}),
		testCase(
			"DELETE should return correct body, headers ans status when id=1",
			test.RestMockTestCaseRequest{
				Type:    "DELETE",
				Path:    "/planets/1",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status: 204,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}),
	}
}

// GetCRUDAPICases Returns all Test Cases
func GetCRUDAPICases(t *testing.T) []test.RestAPITestCase {
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
