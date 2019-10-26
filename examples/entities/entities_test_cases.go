package examplesentities

import (
	"testing"

	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/test"
)

type body struct {
	ID         string   `json:"id" yaml:"id"`
	Name       string   `json:"name" yaml:"name"`
	Type       string   `json:"type" yaml:"type"`
	Period     int      `json:"period" yaml:"period"`
	Atmosphere []string `json:"atmosphere" yaml:"atmosphere"`
}

// GetEntitiesMockCases Returns all Test Cases
func GetEntitiesMockCases(t *testing.T) []test.RestMockTestCase {
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
				BodyFile: "data-all.json",
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
			"GET should return 404 when id=42",
			test.RestMockTestCaseRequest{
				Type:    "GET",
				Path:    "/planets/42",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
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
				BodyFile: "data-new.json",
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
			"PUT should return 404 when id=42",
			test.RestMockTestCaseRequest{
				Type:    "PUT",
				Path:    "/planets/42",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
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
			"DELETE should return 404 when id=42",
			test.RestMockTestCaseRequest{
				Type:    "DELETE",
				Path:    "/planets/42",
				Headers: map[string]string{},
			},
			test.RestMockTestCaseResponse{
				Status:  404,
				Headers: map[string]string{},
				Body:    "",
			}),
	}
}

// GetEntitiesAPICases Returns all Test Cases
func GetEntitiesAPICases(t *testing.T) []test.RestAPITestCase {
	testCase := test.RestAPTestCaseFactory(t)

	return []test.RestAPITestCase{
		testCase(
			&test.RestAPITestCaseConfig{
				Name:             "GET should return all configuration",
				Method:           "GET",
				Status:           200,
				URI:              "/_api/rest/entities",
				RequestFile:      "",
				ResponseFile:     "api-get-all.json",
				ExpectedResponse: []models.EntityRestDto{},
				ActualResponse:   []models.EntityRestDto{},
			}),
	}
}
