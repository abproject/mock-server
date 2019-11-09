package comparator

import (
	"github.com/google/go-cmp/cmp"
)

type ComparatorService struct{}

func NewComparator() Comparator {
	return &ComparatorService{}
}

func (comparator *ComparatorService) Equal(expected interface{}, actual interface{}) bool {
	return cmp.Equal(expected, actual)
}
