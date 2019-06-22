package rest

import (
	"fmt"
	. "github.com/abproject/mock-server/internal_/config"
	"log"
	"net/http"
	"regexp"
	. "strings"
)

type Request struct {
	Method    string
	Path      string
	IsPathReg bool
	Headers   map[string][]string
	source    RestRequestConfig
}

func NewRequest(config RestRequestConfig) *Request {
	request := new(Request)
	request.parse(config)
	return request
}

func (request *Request) Patch(config RestRequestConfig) {
	for headerKey, headers := range config.Headers {
		if _, exist := request.Headers[headerKey]; !exist {
			request.Headers[headerKey] = Split(headers, ";")
		}
	}
}

func (request *Request) IsEqual(r *http.Request) bool {
	if !request.isValid() {
		return false
	}
	if request.Method != "ALL" && request.Method != r.Method {
		return false
	}
	if !request.IsPathReg {
		var path = request.Path
		if path[0] != '/' {
			path = "/" + path
		}
		if path != r.RequestURI {
			return false
		}
	}

	if request.IsPathReg {
		var isMatch, _ = regexp.MatchString(request.Path, r.RequestURI)
		if !isMatch {
			return false
		}
	}

	if len(request.Headers) > 0 {
		var httpHeaders = make(map[string][]string)
		for headerKey, headers := range r.Header {
			httpHeaders[ToUpper(headerKey)] = headers
		}
		for headerKey, headers := range request.Headers {
			for _, headerValue := range headers {
				var key = ToUpper(headerKey)
				if !contains(httpHeaders[key], headerValue) {
					return false
				}
			}
		}
	}

	return true
}

func (request *Request) CompareTo(otherRequest *Request) bool {
	if !otherRequest.isValid() {
		return true
	} else if !request.isValid() {
		return false
	}

	if request.Method != "ALL" && otherRequest.Method == "ALL" {
		return true
	} else if request.Method == "ALL" && otherRequest.Method != "ALL" {
		return false
	} else {
		var requestHeadersAmount = 0
		for _, headers := range request.Headers {
			requestHeadersAmount += len(headers)
		}
		var otherRequestHeadersAmount = 0
		for _, headers := range otherRequest.Headers {
			otherRequestHeadersAmount += len(headers)
		}

		if requestHeadersAmount > otherRequestHeadersAmount {
			return true
		} else if requestHeadersAmount < otherRequestHeadersAmount {
			return false
		} else {
			if !request.IsPathReg {
				return true
			} else if !otherRequest.IsPathReg {
				return false
			}
		}
	}
	return true
}

func (request *Request) parse(config RestRequestConfig) {
	if config.Path == "" && config.PathReg == "" {
		log.Fatal(fmt.Sprintf("Rest config: request 'path' or 'pathReg' is required\n%#v", config))
	}

	var method = ToUpper(config.Method)
	if method == "" {
		method = "ALL"
	}
	var isRegPath bool
	var path string

	if Contains(config.Path, "/:") {
		isRegPath = true
		r, _ := regexp.Compile("(:[a-zA-Z0-9_-]+)")
		path = "^/" + r.ReplaceAllString(config.Path, "[a-zA-Z0-9_-]+") + "$"
	} else {
		isRegPath = config.PathReg != ""
		path = config.Path
		if isRegPath {
			path = config.PathReg
		}
	}

	var headers = make(map[string][]string)
	for headerKey, headerValue := range config.Headers {
		headers[headerKey] = Split(headerValue, ";")
	}

	*request = Request{
		method,
		path,
		isRegPath,
		headers,
		config,
	}
}

func (request *Request) isValid() bool {
	if request.Method == "" {
		return false
	}
	if request.Path == "" {
		return false
	}
	return true
}

func contains(array []string, value string) bool {
	for _, n := range array {
		if value == n {
			return true
		}
	}
	return false
}
