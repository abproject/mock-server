package comparator

type Comparator interface {
	Equal(expected interface{}, actual interface{}) bool
	// EqualJSON(expected []byte, actual []byte)
}
