package rest

import (
	. "mock-server/internal/testing"
	"github.com/abproject/mock-server/internal/testing"
)

func TestRequestInit(t *testing.T) {
	v := Validation{T: t}

	tables := []struct {
		given        RequestConfig
		expected     Request
		errorMessage string
	}{
		{
			RequestConfig{},
			Request{},
			"Request 'path' or 'pathReg' is required\n",
		},
		{
			RequestConfig{
				Path: "path/test",
			},
			Request{
				Method:    "ALL",
				Path:      "path/test",
				IsPathReg: false,
				Headers: map[string][]string{},
				source: RequestConfig{
					Path: "path/test",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method: "GET",
				Path:   "path",
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method: "post",
				Path:   "path",
			},
			Request{
				Method:    "POST",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{},
				source: RequestConfig{
					Method: "post",
					Path:   "path",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method: "GET",
				Path:   "path/:id",
			},
			Request{
				Method:    "GET",
				Path:      "^/path/[a-zA-Z0-9_-]+$",
				IsPathReg: true,
				Headers: map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path/:id",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method: "GET",
				Path:   "path/:id/:code",
			},
			Request{
				Method:    "GET",
				Path:      "^/path/[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$",
				IsPathReg: true,
				Headers: map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path/:id/:code",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method:  "GET",
				PathReg: "path?",
			},
			Request{
				Method:    "GET",
				Path:      "path?",
				IsPathReg: true,
				Headers: map[string][]string{},
				source: RequestConfig{
					Method:  "GET",
					PathReg: "path?",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method:  "GET",
				Path:    "path",
				PathReg: "path?",
			},
			Request{
				Method:    "GET",
				Path:      "path?",
				IsPathReg: true,
				Headers: map[string][]string{},
				source: RequestConfig{
					Method:  "GET",
					Path:    "path",
					PathReg: "path?",
				},
			},
			"",
		},
		{
			RequestConfig{
				Method: "GET",
				Path:   "test",
				Headers: map[string]string{
					"header1": "value1;value2",
					"header2": "value3",
				},
			},
			Request{
				Method: "GET",
				Path:   "test",
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
				source: RequestConfig{
					Method: "GET",
					Path:   "test",
					Headers: map[string]string{
						"header1": "value1;value2",
						"header2": "value3",
					},
				},
			},
			"",
		},
	}

	for idx, table := range tables {
		v.SetTestCase(idx, table)
		var requestConfig = table.given
		var expectedRequest = table.expected
		var expectedErrorMessage = table.errorMessage

		var request Request
		var err = request.Init(requestConfig)

		v.IsEqual(ValidationConfig{
			Expected: expectedRequest,
			Given:    request,
		})
		v.IsEqualError(ValidationConfig{
			Expected: expectedErrorMessage,
			Given:    err,
		})
	}
}
