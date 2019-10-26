package rest

import (
	"encoding/json"
	"testing"

	"github.com/abproject/mock-server/internal/models"
)

var compareTestCases = []struct {
	Request1 models.RequestRestDto `json:"request1"`
	Request2 models.RequestRestDto `json:"request2"`
	expect   bool
}{
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		false,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path2",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path2",
			PathReg: "",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		false,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: nil,
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{},
		},
		true,
	},
	{
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "",
			PathReg: "path",
			Headers: map[string]string{
				"header1": "value1;value2",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
		models.RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header1": "value1",
				"header2": "value3",
			},
		},
		models.RequestRestDto{
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
