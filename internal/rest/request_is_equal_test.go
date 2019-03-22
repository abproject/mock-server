package rest

import (
	. "mock-server/internal/testing"
	"net/http"
	"github.com/abproject/mock-server/internal/testing"
)
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
