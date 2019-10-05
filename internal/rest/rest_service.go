package rest

import (
	"net/http"
	"regexp"
	"strings"
)

// IsEqual compares Storage Entity with http.Request
func IsEqual(entity entityRest, r *http.Request) bool {
	request := entity.Config.Request

	method := strings.ToUpper(request.Method)
	if method != "" && method != strings.ToUpper(r.Method) {
		return false
	}

	if request.PathReg != "" {
		var isMatch, _ = regexp.MatchString(request.PathReg, r.RequestURI)
		if !isMatch {
			return false
		}
	} else {
		var path = request.Path
		if len(path) == 0 || path[0] != '/' {
			path = "/" + path
		}
		regDynamicID := regexp.MustCompile(`:[A-Za-z0-9_-]+`)
		regPathString := regDynamicID.ReplaceAllString(path, "[A-Za-z0-9_-]+")
		regPath := regexp.MustCompile("(?i)^" + regPathString + "$")
		if !regPath.MatchString(r.RequestURI) {
			return false
		}
	}

	if len(request.Headers) > 0 {
		var httpHeaders = make(map[string][]string)
		for headerKey, headers := range r.Header {
			httpHeaders[strings.ToUpper(headerKey)] = headers
		}

		for headerKey, headers := range request.Headers {
			splitHeaders := strings.Split(headers, ";")
			for _, headerValue := range splitHeaders {
				var key = strings.ToUpper(headerKey)
				if !contains(httpHeaders[key], headerValue) {
					return false
				}
			}
		}
	}

	return true
}

// Compare requests for sorting
func Compare(request1 RequestRestDto, request2 RequestRestDto) bool {
	if request1.Method != "" && request2.Method == "" {
		return true
	} else if request1.Method == "" && request2.Method != "" {
		return false
	}

	var requestHeadersAmount = 0
	for _, headers := range request1.Headers {
		requestHeadersAmount += len(strings.Split(headers, ";"))
	}
	var otherRequestHeadersAmount = 0
	for _, headers := range request2.Headers {
		otherRequestHeadersAmount += len(strings.Split(headers, ";"))
	}

	if requestHeadersAmount > otherRequestHeadersAmount {
		return true
	} else if requestHeadersAmount < otherRequestHeadersAmount {
		return false
	}

	if request1.PathReg == "" {
		return true
	} else if request2.PathReg == "" {
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
