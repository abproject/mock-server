package testprinter

import (
	"encoding/json"
	"runtime/debug"
	"strings"
	"testing"
)

type Printer interface {
	ComparationError(expected interface{}, actual interface{}, message string)
}

type TestPrinter struct {
	t     *testing.T
	title string
}

func NewTestPrinter(t *testing.T, title string) Printer {
	return &TestPrinter{
		t:     t,
		title: title,
	}
}

func (printer *TestPrinter) ComparationError(expected interface{}, actual interface{}, message string) {
	frame := strings.Repeat("*", 6+len(message))
	printer.t.Errorf(
		`

TESTCASE: %s

%s
*  %s  *
%s

Expected:
%s
Actual:
%s
`, printer.title, frame, message, frame,
		prettyPrint(expected),
		prettyPrint(actual))
	printer.t.Log(string(debug.Stack()))
}

func prettyPrint(instance interface{}) string {
	bytes, ok := instance.([]byte)
	if ok {
		return string(bytes)
	}
	formattedValue, _ := json.MarshalIndent(instance, "", "  ")
	return string(formattedValue)
}
