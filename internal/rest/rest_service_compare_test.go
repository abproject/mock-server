package rest

import (
	"encoding/json"
	"testing"
)

var compareTestCases = []struct {
	Request1 RequestRestDto `json:"request1"`
	Request2 RequestRestDto `json:"request2"`
	expect   bool
}{
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path2",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path2",
			PathReg: "",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{},
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		false,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		true,
	},
	{
		RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		false,
	},
}

func TestRestServiceCompare(t *testing.T) {
	counter := 0
	for i, testCase := range compareTestCases {
		actual := Compare(testCase.Request1, testCase.Request2)
		if testCase.expect != actual {
			json, err := json.MarshalIndent(testCase, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("Compare Test Case #%d:\n%s\nExpect: %t\nActual:  %t", i, string(json), testCase.expect, actual)
			counter++
		}
	}
	if counter > 0 {
		t.Errorf("Compare Failed/Total: %d/%d", counter, len(compareTestCases))
	}
}
