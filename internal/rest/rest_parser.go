package rest

import (
	"sync"
)

var instanceParser iRestParser
var onceParser sync.Once

type iRestParser interface {
}

type restParser struct {
	//storage iRestStorage
}

func GetParser() iRestParser {
	onceService.Do(func() {
		instanceParser = newParser()
	})
	return instanceService
}

func newParser() iRestParser {
	return &restParser{}
}
