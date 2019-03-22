package rest

import (
	. "mock-server/internal/testing"
	"testing"
)

func TestResponseInit(t *testing.T) {
	v := Validation{T: t}

	tables := []struct {
		given        ResponseConfig
		expected     Response
	}{
		{
			ResponseConfig{},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
		},
		{
			ResponseConfig{
				Body: "Body",
			},
			Response{
				Body: "Body",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{
					Body: "Body",
				},
			},
		},
		{
			ResponseConfig{
				Status: 201,
			},
			Response{
				Body: "",
				Status: 201,
				Headers: map[string]string{},
				source: ResponseConfig{
					Status: 201,
				},
			},
		},
		{
			ResponseConfig{
				Status: 0,
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{
					Status: 0,
				},
			},
		},
		{
			ResponseConfig{
				Headers:  map[string]string{
					"header1": "value1;value2",
					"header2": "value3",
				},
			},
			Response{
				Body: "",
				Status: 200,
				Headers:  map[string]string{
					"header1": "value1;value2",
					"header2": "value3",
				},
				source: ResponseConfig{
					Headers:  map[string]string{
						"header1": "value1;value2",
						"header2": "value3",
					},
				},
			},
		},
	}

	for idx, table := range tables {
		v.SetTestCase(idx, table)
		var responseConfig = table.given
		var expectedResponse = table.expected

		var response Response
		response.Init(responseConfig)

		v.IsEqual(ValidationConfig{
			Expected: expectedResponse,
			Given:    response,
		})
	}
}
