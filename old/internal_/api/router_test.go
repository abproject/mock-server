package api

import (
	"reflect"
	"testing"
)

func TestGetPathVariables(t *testing.T) {
	type TestCase struct {
		path       string
		requestUri string
		vars       map[string]string
	}
	tests := []TestCase{
		{
			path:       "/_api/rest",
			requestUri: "/_api/rest",
			vars: make(map[string]string),
		},
		{
			path:       "/_api/rest/:id",
			requestUri: "/_api/rest/1234",
			vars: map[string]string {
				"id": "1234",
			},
		},
		{
			path:       "/_api/rest/:id/:name",
			requestUri: "/_api/rest/1234/abc",
			vars: map[string]string {
				"id": "1234",
				"name": "abc",
			},
		},
	}

	for _, test := range tests {
		vars := getPathVariables(test.path, test.requestUri)
		if !reflect.DeepEqual(vars, test.vars) {
			t.Errorf(`
Failed TestGetPathVariables:
endpoint: '%#v'
requestUri:	'%s'
	%#v expected to be %#v`,
				test.path,
				test.requestUri,
				vars,
				test.vars)
		}
	}
}
