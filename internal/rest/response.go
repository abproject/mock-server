package rest

import (
	"fmt"
	. "github.com/abproject/mock-server/internal/config"
	"io/ioutil"
	"log"
)

type Response struct {
	Body    string
	File    []byte
	Status  int
	Headers map[string]string
	source  RestResponseConfig
}

func NewResponse(config RestResponseConfig) *Response {
	var status = config.Status
	if status == 0 {
		status = 200
	}

	var headers = make(map[string]string)
	for headerKey, headerValue := range config.Headers {
		headers[headerKey] = headerValue
	}

	var file []byte
	if config.File != "" {
		data, err := ioutil.ReadFile(config.File)
		if err != nil {
			log.Fatal(fmt.Sprintf("Rest Config: Required file in config '%s' not found", config.File))
		}
		file = data
	}

	return &Response{
		Body:    config.Body,
		File:    file,
		Status:  status,
		Headers: headers,
		source:  config,
	}
}

func (response *Response) Patch(config RestResponseConfig) {

	if response.Body == "" && config.Body != "" {
		response.Body = config.Body
	}

	if response.source.Status == 0 && config.Status != 0 {
		response.Status = config.Status
	}

	for headerKey, headers := range config.Headers {
		if _, exist := response.Headers[headerKey]; !exist {
			response.Headers[headerKey] = headers
		}
	}
}
