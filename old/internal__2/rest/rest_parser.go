package rest

import (
	"sync"
)

var instanceParser RestParser
var onceParser sync.Once

type RestParser interface {
}

type restParser struct {
	//storage RestStorage
}

func GetParser() RestParser {
	onceService.Do(func() {
		instanceParser = newParser()
	})
	return instanceService
}

func newParser() RestParser {
	return &restParser{}
}
