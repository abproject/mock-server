package api

import (
	"log"
	"net/http"
	"regexp"
	. "strings"
)

type HandlerData struct {
	w http.ResponseWriter
	r *http.Request
	vars map[string]string
}

type Handler func(HandlerData)

type endpoint struct {
	request string
	path string
	regPath    regexp.Regexp
	handler Handler
}

var endpoints []endpoint

type Router struct{}

func (router *Router) Get(path string, handler Handler) {
	endpoints = append(endpoints, endpoint{"GET", path, parsePath(path), handler})
}

func (router *Router) Post(path string, handler Handler) {
	endpoints = append(endpoints, endpoint{"POST", path,parsePath(path), handler})
}

func (router *Router) Put(path string, handler Handler) {
	endpoints = append(endpoints, endpoint{"PUT", path,parsePath(path), handler})
}

func (router *Router) Delete(path string, handler Handler) {
	endpoints = append(endpoints, endpoint{"DELETE", path, parsePath(path), handler})
}

func parsePath(path string) regexp.Regexp {
	dynamic, _ := regexp.Compile("(:[a-zA-Z0-9_-]+)")
	r, _ := regexp.Compile("^" + dynamic.ReplaceAllString(path, "([a-zA-Z0-9_-]+)") + "$")
	return *r
}

func (router *Router) route(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	for _, endpoint := range endpoints {
		if ToUpper(endpoint.request) == ToUpper(r.Method) && endpoint.regPath.MatchString(r.RequestURI) {
			log.Printf("%s: %#v", r.RequestURI, r.Method)
			endpoint.handler(HandlerData{w,r, getPathVariables(endpoint.path, r.RequestURI)})
			return
		}
	}
	// No endpoint found
	log.Printf("%s: %v", r.RequestURI, "No API Endpoint found")
	w.WriteHeader(http.StatusNotFound)
}

func getPathVariables(path string, requestUri string) map[string]string {
	vars := make(map[string]string)
	r := parsePath(path)
	dynamic, _ := regexp.Compile(":[a-zA-Z0-9_-]+")
	names := dynamic.FindAllString(path, -1)
	values := r.FindStringSubmatch(requestUri)
	for i, name := range names {
		n := Replace(name, ":", "", -1)
		vars[n] = values[i + 1]
	}
	return vars
}