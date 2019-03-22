package rest

import (
	. "mock-server/internal/testing"
	"testing"
)

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
