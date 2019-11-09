package comparator

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
)

type ComparatorService struct{}

func NewComparator() Comparator {
	return &ComparatorService{}
}

func (comparator *ComparatorService) Equal(expected interface{}, actual interface{}) bool {
	return cmp.Equal(expected, actual)
}

func (comparator *ComparatorService) EqualJson(expected []byte, actual []byte) bool {
	if cmp.Equal(expected, actual) {
		return true
	}
	var expectedInstance, actualInstance interface{}
	if err := json.Unmarshal(expected, &expectedInstance); err != nil {
		return false
	}
	if err := json.Unmarshal(actual, &actualInstance); err != nil {
		return false
	}
	return cmp.Equal(expectedInstance, actualInstance)
}
