package rest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/abproject/mock-server/internal/models"
)

type Request struct {
	Method     string
	RequestURI string
	Header     map[string][]string
}

var isEqualTestCases = []struct {
	Dto     models.RequestRestDto `json:"entity"`
	Request Request                   `json:"request"`
	expect  bool
}{
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "",
		},
		Request{
			Method:     "GET",
			RequestURI: "/",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "get",
			Path:   "",
		},
		Request{
			Method:     "GET",
			RequestURI: "/",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "Get",
			Path:   "",
		},
		Request{
			Method:     "GET",
			RequestURI: "/",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "",
			Path:   "path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "",
			Path:   "path",
		},
		Request{
			Method:     "POST",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "/",
		},
		Request{
			Method:     "GET",
			RequestURI: "/",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/PATH",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "PATH",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "/path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
		},
		Request{
			Method:     "POST",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		false,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path/hello/world",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path/hello/world",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "//path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		false,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "/path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/pat",
			Header:     map[string][]string{},
		},
		false,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "^/path/.*",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path/hello",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "^/path/.*",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		false,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "^/path(/.*)?",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path/hello",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "^/path(/.*)?",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
		},
		Request{
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
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header1": "header1-value1;header1-value2",
				"header2": "header2-value",
			},
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header:     map[string][]string{},
		},
		false,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header1": "header1-value1;header1-value2",
				"header2": "header2-value",
			},
		},
		Request{
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
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header1": "header1-value1;header1-value2",
				"header2": "header2-value",
			},
		},
		Request{
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
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header1": "header1-value1;header1-value2",
				"header2": "header2-value",
			},
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header: map[string][]string{
				"header1": {"header1-value1", "header1-value2"},
			},
		},
		false,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header1": "header1-value1;header1-value3",
				"header2": "header2-value",
			},
		},
		Request{
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
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header1": "header1-value2;header1-value1;header1-value-3",
			},
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header: map[string][]string{
				"header1": {"header1-value-3", "header1-value1", "header1-value2"},
			},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
			Headers: map[string]string{
				"header-name": "header1-value2;header1-value1;header1-value-3",
			},
		},
		Request{
			Method:     "GET",
			RequestURI: "/path",
			Header: map[string][]string{
				"Header-Name": {"header1-value-3", "header1-value1", "header1-value2"},
			},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path/:id",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path/1",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path/:id",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path/42",
			Header:     map[string][]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method: "GET",
			Path:   "path",
		},
		Request{
			Method:     "GET",
			RequestURI: "/path/1",
			Header:     map[string][]string{},
		},
		false,
	},
}

func TestRestServiceIsEqual(t *testing.T) {
	counter := 0
	for i, testCase := range isEqualTestCases {
		request := &http.Request{
			Method:     testCase.Request.Method,
			RequestURI: testCase.Request.RequestURI,
			Header:     testCase.Request.Header,
		}
		entityRest := entityRest{
			Config: models.EndpointRestDto{
				Request: models.RequestRestDto{
					Method:  testCase.Dto.Method,
					Path:    testCase.Dto.Path,
					PathReg: testCase.Dto.PathReg,
					Headers: testCase.Dto.Headers,
				},
			},
		}
		actual := IsEqual(entityRest, request)
		if testCase.expect != actual {
			json, err := json.MarshalIndent(testCase, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("IsEqual Test Case #%d:\n%s\nExpect: %t\nActual:  %t", i, string(json), testCase.expect, actual)
			counter++
		}
	}
	if counter > 0 {
		t.Errorf("IsEqual Failed/Total: %d/%d", counter, len(isEqualTestCases))
	}
}
