package sharedcomparator

type Comparator interface {
	Equal(expected interface{}, actual interface{}, message string)
}
