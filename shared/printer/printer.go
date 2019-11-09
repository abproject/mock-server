package sharedprinter

type IPrinter interface {
	Error(expected interface{}, actual interface{}, message string)
}
