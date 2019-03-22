package rest

import (
	. "mock-server/internal/testing"
	"github.com/abproject/mock-server/internal/testing"
)

func TestResponsePatch(t *testing.T) {
	v := Validation{T: t}

	tables := []struct {
		givenConfig   ResponseConfig
		givenResponse Response
		expected      Response
	}{
		{
			ResponseConfig{},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
		},
		{
			ResponseConfig{},
			Response{
				Body: "",
				Status: 201,
				Headers: map[string]string{},
				source: ResponseConfig{
					Status: 201,
				},
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
				Status: 300,
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
			Response{
				Body: "",
				Status: 300,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
		},
		{
			ResponseConfig{
				Status: 300,
			},
			Response{
				Body: "",
				Status: 201,
				Headers: map[string]string{},
				source: ResponseConfig{
					Status: 201,
				},
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
				Body: "config-body",
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
			Response{
				Body: "config-body",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
		},
		{
			ResponseConfig{
				Body: "config-body",
			},
			Response{
				Body: "response-body",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{
					Body: "response-body",
				},
			},
			Response{
				Body: "response-body",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{
					Body: "response-body",
				},
			},
		},
		{
			ResponseConfig{},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{
					"header-1": "header-value1;header-value2",
					"header-2": "header-value3;header-value4",
				},
				source: ResponseConfig{
					Headers: map[string]string{
						"header-1": "header-value1;header-value2",
						"header-2": "header-value3;header-value4",
					},
				},
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{
					"header-1": "header-value1;header-value2",
					"header-2": "header-value3;header-value4",
				},
				source: ResponseConfig{
					Headers: map[string]string{
						"header-1": "header-value1;header-value2",
						"header-2": "header-value3;header-value4",
					},
				},
			},
		},
		{
			ResponseConfig{
				Headers: map[string]string{
					"header-1": "header-value1;header-value2",
					"header-2": "header-value3;header-value4",
				},
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{},
				source: ResponseConfig{},
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{
					"header-1": "header-value1;header-value2",
					"header-2": "header-value3;header-value4",
				},
				source: ResponseConfig{},
			},
		},

		{
			ResponseConfig{
				Headers: map[string]string{
					"header-1": "header-value5;header-value6",
					"header-3": "header-value-7",
				},
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{
					"header-1": "header-value1;header-value2",
					"header-2": "header-value3;header-value4",
				},
				source: ResponseConfig{
					Headers: map[string]string{
						"header-1": "header-value1;header-value2",
						"header-2": "header-value3;header-value4",
					},
				},
			},
			Response{
				Body: "",
				Status: 200,
				Headers: map[string]string{
					"header-1": "header-value1;header-value2",
					"header-2": "header-value3;header-value4",
					"header-3": "header-value-7",
				},
				source: ResponseConfig{
					Headers: map[string]string{
						"header-1": "header-value1;header-value2",
						"header-2": "header-value3;header-value4",
					},
				},
			},
		},
	}

	for idx, table := range tables {
		v.SetTestCase(idx, table)
		var responseConfig = table.givenConfig
		var givenResponse = table.givenResponse
		var expectedResponse = table.expected

		givenResponse.Patch(responseConfig)

		v.IsEqual(ValidationConfig{
			Expected: expectedResponse,
			Given:    givenResponse,
		})
	}
}
