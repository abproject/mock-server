package rest

import (
	. "mock-server/internal/testing"
	"github.com/abproject/mock-server/internal/testing"
)

func TestRequestPatch(t *testing.T) {
	v := Validation{T: t}

	tables := []struct {
		givenConfig  RequestConfig
		givenRequest Request
		expected     Request
	}{
		{
			RequestConfig{},
			Request{},
			Request{},
		},
		{
			RequestConfig{},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
		},
		{
			RequestConfig{},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
		},
		{
			RequestConfig{},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
		},
		{
			RequestConfig{
				Headers: map[string]string{
					"header1": "header1-value1;header1-value3",
					"header2": "header2-value",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
		},
		{
			RequestConfig{
				Headers: map[string]string{
					"header1": "header1-value3;header1-value1",
					"header2": "header2-value",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
		},
		{
			RequestConfig{
				Headers: map[string]string{
					"header1": "header1-value3;header1-value1",
					"header2": "header2-value",
					"header4": "header4-value",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
					"header4": {"header4-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
		},
		{
			RequestConfig{
				Headers: map[string]string{
					"header1": "header1-value3;header1-value1",
					"header2": "header2-new-value",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1;header1-value3",
						"header2": "header2-value",
					},
				},
			},
		},
		{
			RequestConfig{
				Headers: map[string]string{
					"header1": "header1-value3;header1-value1",
					"header2": "header2-new-value",
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1",
						"header2": "header2-value",
					},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1"},
					"header2": {"header2-value"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
					Headers: map[string]string{
						"header1": "header1-value1",
						"header2": "header2-value",
					},
				},
			},
		},
	}

	for idx, table := range tables {
		v.SetTestCase(idx, table)
		var requestConfig = table.givenConfig
		var currentRequest = table.givenRequest
		var expectedRequest = table.expected

		currentRequest.Patch(requestConfig)

		v.IsEqual(ValidationConfig{
			Expected: expectedRequest,
			Given:    currentRequest,
		})
	}
}
