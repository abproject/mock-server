package rest

import (
	. "github.com/abproject/mock-server/validation"
	"net/http"
	"testing"
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

func TestRequestIsEqual(t *testing.T) {
	v := Validation{T: t}

	tables := []struct {
		request     Request
		httpRequest http.Request
		isEqual     bool
	}{
		{
			Request{},
			http.Request{},
			false,
		},
		{
			Request{
				Method: "GET",
				Path: "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			Request{
				Method: "POST",
				Path: "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			Request{
				Method: "ALL",
				Path: "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},

			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "path/hello/world",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path/hello/world",
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "//path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/pat",
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "pat",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "^/path/.*",
				IsPathReg: true,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path/hello",
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "^/path/.*",
				IsPathReg: true,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "^/path(/.*)?",
				IsPathReg: true,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path/hello",
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "^/path(/.*)?",
				IsPathReg: true,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{},
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method: "GET",
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
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
				},
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value1", "header1-value3"},
					"header2": {"header2-value"},
				},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value1", "header1-value2"},
					"header2": {"header2-value"},
				},
			},
			false,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"header1-value2", "header1-value1", "header1-value-3"},
				},
			},
			http.Request{
				Method: "GET",
				RequestURI: "/path",
				Header: map[string][]string{
					"header1": {"header1-value-3", "header1-value1", "header1-value2"},
				},
			},
			true,
		},
		{
			Request{
				Method: "GET",
				Path:      "/path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header-name": {"header1-value2", "header1-value1", "header1-value-3"},
				},
			},
			http.Request{
				Method: "GET",
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

		v.IsEqual(ValidationConfig{
			Expected: ExpectedIsEqual,
			Given:    equal,
		})
	}
}


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


func TestRequestCompareTo(t *testing.T) {
	v := Validation{T: t}

	tables := []struct {
		request1 Request
		request2 Request
		compare  bool
	}{
		{
			Request{},
			Request{},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			false,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path2",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path2",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			false,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers:   map[string][]string{},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers: map[string][]string{},
			},
			Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			true,
		},
		{
			Request{
				Method:    "ALL",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{},
			},
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers: map[string][]string{},
			},

			false,
		},
		{
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: true,
				Headers: map[string][]string{
					"header1": {"value1", "value2"},
					"header2": {"value3"},
				},
			},
			Request{
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
			Request{
				Method:    "GET",
				Path:      "path",
				IsPathReg: false,
				Headers: map[string][]string{
					"header1": {"value1"},
					"header2": {"value3"},
				},
			},
			Request{
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

		v.IsEqual(ValidationConfig{
			Expected: expectedCompare,
			Given:    compare,
		})
	}
}
