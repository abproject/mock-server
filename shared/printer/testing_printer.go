package sharedprinter

import (
	"encoding/json"
	"runtime/debug"
	"strings"
	"testing"
)

type TestingPrinter struct {
	t     *testing.T
	title string
}

func NewTestingPrinter(t *testing.T, title string) IPrinter {
	return &TestingPrinter{
		t:     t,
		title: title,
	}
}

func (printer *TestingPrinter) Error(expected interface{}, actual interface{}, message string) {
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

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
