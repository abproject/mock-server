package comparator

import (
	"testing"

	"github.com/abproject/mock-server/shared/testprinter"
)

type TestStructExample struct {
	A string
	B int
	C bool
}

func TestBaseComparatorEqual(t *testing.T) {
	comparator := NewComparator()
	testCases := []struct {
		printer testprinter.Printer
		param1  interface{}
		param2  interface{}
		isEqual bool
	}{
		{
			printer: testprinter.NewTestPrinter(t, "Integers equal"),
			param1:  42,
			param2:  42,
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Integers not equal"),
			param1:  42,
			param2:  13,
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Integers not equal: nil"),
			param1:  42,
			param2:  nil,
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Strings equal"),
			param1:  "str",
			param2:  "str",
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Strings equal: empty"),
			param1:  "",
			param2:  "",
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Strings not equal: substring"),
			param1:  "str",
			param2:  "str1",
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Strings not equal: nil"),
			param1:  "str",
			param2:  nil,
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps equal"),
			param1: map[string]string{
				"key": "value",
			},
			param2: map[string]string{
				"key": "value",
			},
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps equal: both empty"),
			param1:  map[string]string{},
			param2:  map[string]string{},
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps not equal: key is substring"),
			param1: map[string]string{
				"key1": "value",
			},
			param2: map[string]string{
				"key": "value",
			},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps not equal: value is substring"),
			param1: map[string]string{
				"key": "value",
			},
			param2: map[string]string{
				"key": "value1",
			},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps not equal: one empty"),
			param1: map[string]string{
				"key": "value",
			},
			param2:  map[string]string{},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps not equal: more entries"),
			param1: map[string]string{
				"key": "value",
			},
			param2: map[string]string{
				"key":  "value",
				"key2": "value2",
			},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Maps not equal: nil"),
			param1: map[string]string{
				"key": "value",
			},
			param2:  nil,
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Arrays equal"),
			param1:  []string{"str", "str2"},
			param2:  []string{"str", "str2"},
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Arrays equal: both empty"),
			param1:  []string{},
			param2:  []string{},
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Arrays not equal: value is substring"),
			param1:  []string{"str", "str2"},
			param2:  []string{"str1", "str2"},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Arrays not equal: one empty"),
			param1:  []string{"str", "str2"},
			param2:  []string{},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Arrays not equal: more entries"),
			param1:  []string{"str", "str2"},
			param2:  []string{"str"},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Arrays not equal: nil"),
			param1:  []string{"str", "str2"},
			param2:  nil,
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Structs equal"),
			param1: TestStructExample{
				A: "str",
				B: 42,
				C: true,
			},
			param2: TestStructExample{
				A: "str",
				B: 42,
				C: true,
			},
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Structs equal: both empty"),
			param1:  TestStructExample{},
			param2:  TestStructExample{},
			isEqual: true,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Structs not equal: value is substring"),
			param1: TestStructExample{
				A: "str",
				B: 42,
				C: true,
			},
			param2: TestStructExample{
				A: "str1",
				B: 42,
				C: true,
			},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Structs not equal: one empty"),
			param1: TestStructExample{
				A: "str",
				B: 42,
				C: true,
			},
			param2:  TestStructExample{},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Structs not equal: one is default"),
			param1: TestStructExample{
				A: "str",
				B: 42,
				C: true,
			},
			param2: TestStructExample{
				A: "str",
				C: true,
			},
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "Structs not equal: nil"),
			param1: TestStructExample{
				A: "str",
				B: 42,
				C: true,
			},
			param2:  nil,
			isEqual: false,
		},
		{
			printer: testprinter.NewTestPrinter(t, "nil are equals"),
			param1:  nil,
			param2:  nil,
			isEqual: true,
		},
	}

	for _, testCase := range testCases {
		if testCase.isEqual != comparator.Equal(testCase.param1, testCase.param2) {
			var message string
			if testCase.isEqual == true {
				message = "must be equal"
			} else {
				message = "must not be equal"
			}
			testCase.printer.ComparationError(testCase.param1, testCase.param2, message)
		}
	}
}
