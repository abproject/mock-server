package comparator

type Comparator interface {
	Equal(expected interface{}, actual interface{}) bool
	EqualJson(expected []byte, actual []byte) bool
}
