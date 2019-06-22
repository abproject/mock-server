package rest

import (
	"github.com/abproject/mock-server/internal_"
	"net/http"
	"testing"
)

func TestRequestInit(t *testing.T) {
	v := internal.Validation{T: t}

	tables := []struct {
		given        RequestConfig
		expected     internal.Request
		errorMessage string
	}{
		{
			RequestConfig{},
			internal.Request{},
			"request 'path' or 'pathReg' is required\n",
		},
		{
			RequestConfig{
				Path: "path/test",
			},
			internal.Request{
				Method:    "ALL",
				Path:      "path/test",
				IsPathReg: false,
				Headers:   map[string][]string{},
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
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
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
			internal.Request{
				Method:    "POST",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
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
			internal.Request{
				Method:    "GET",
				Path:      "^/path/[a-zA-Z0-9_-]+$",
				IsPathReg: true,
				Headers:   map[string][]string{},
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
			internal.Request{
				Method:    "GET",
				Path:      "^/path/[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$",
				IsPathReg: true,
				Headers:   map[string][]string{},
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
			internal.Request{
				Method:    "GET",
				Path:      "path?",
				IsPathReg: true,
				Headers:   map[string][]string{},
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
			internal.Request{
				Method:    "GET",
				Path:      "path?",
				IsPathReg: true,
				Headers:   map[string][]string{},
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
			internal.Request{
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

		var request internal.Request
		var err = request.Init(requestConfig)

		v.IsEqual(internal.ValidationConfig{
			Expected: expectedRequest,
			Given:    request,
		})
		v.IsEqualError(internal.ValidationConfig{
			Expected: expectedErrorMessage,
			Given:    err,
		})
	}
}

func TestRequestIsEqual(t *testing.T) {
	v := internal.Validation{T: t}

	tables := []struct {
		request     internal.Request
		httpRequest http.Request
		isEqual     bool
	}{
		{
			internal.Request{},
			http.Request{},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			internal.Request{
				Method:    "POST",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},

			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path/hello/world",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path/hello/world",
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "//path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/pat",
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "pat",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "^/path/.*",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path/hello",
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "^/path/.*",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "^/path(/.*)?",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path/hello",
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "^/path(/.*)?",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header:     map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header:     map[string][]string{},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
					"header3": {"header3-value"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
				},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value2", "header1-value1", "header1-value-3"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value-3", "header1-value1", "header1-value2"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header-name": {"header1-value2", "header1-value1", "header1-value-3"},
				},
			},
			http.Request{
				Method:     "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"Header-Name": {"header1-value-3", "header1-value1", "header1-value2"},
				},
			},
			true,
		},
	}

	for idx, table := range tables {
		v.SetTestCase(idx, table)
		var request = table.request
		var httpRequest = table.httpRequest
		var ExpectedIsEqual = table.isEqual

		var equal = request.IsEqual(&httpRequest)

		v.IsEqual(internal.ValidationConfig{
			Expected: ExpectedIsEqual,
			Given:    equal,
		})
	}
}

func TestRequestPatch(t *testing.T) {
	v := internal.Validation{T: t}

	tables := []struct {
		givenConfig  RequestConfig
		givenRequest internal.Request
		expected     internal.Request
	}{
		{
			RequestConfig{},
			internal.Request{},
			internal.Request{},
		},
		{
			RequestConfig{},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
			internal.Request{
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
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
				source: RequestConfig{
					Method: "GET",
					Path:   "path",
				},
			},
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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
			internal.Request{
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

		v.IsEqual(internal.ValidationConfig{
			Expected: expectedRequest,
			Given:    currentRequest,
		})
	}
}

func TestRequestCompareTo(t *testing.T) {
	v := internal.Validation{T: t}

	tables := []struct {
		request1 internal.Request
		request2 internal.Request
		compare  bool
	}{
		{
			internal.Request{},
			internal.Request{},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path2",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path2",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			internal.Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},

			false,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			true,
		},
		{
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			internal.Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			false,
		},
	}

	for idx, table := range tables {
		v.SetTestCase(idx, table)
		var request1 = table.request1
		var request2 = table.request2
		var expectedCompare = table.compare

		var compare = request1.CompareTo(&request2)

		v.IsEqual(internal.ValidationConfig{
			Expected: expectedCompare,
			Given:    compare,
		})
	}
}
