package sharedcomparator

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	sharedprinter "github.com/abproject/mock-server/shared/printer"
)

type TestingComparator struct {
	t       *testing.T
	printer sharedprinter.IPrinter
}

func NewTestingComparator(t *testing.T, title string) Comparator {
	return &TestingComparator{
		t:       t,
		printer: sharedprinter.NewTestingPrinter(t, title),
	}
}

func (comparator *TestingComparator) Equal(expected interface{}, actual interface{}, message string) {
	if !cmp.Equal(expected, actual) {
		comparator.printer.Error(expected, actual, message)
	}
}
