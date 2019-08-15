package router

import (
	"log"
	"net/http"
)

// RouteMock Rest API
func RouteMock(c Context, w http.ResponseWriter, r *http.Request) {
	endpoint, err := (*c.RestStorage).FindByRequest(r)
	if err != nil {
		notFoundMockHandler(w, r)
		return
	}
	// Rest endpoint found
	c.Logger.Printf("%s: \n%+v", r.RequestURI, endpoint)
	for headerKey, headerValue := range endpoint.Response.Headers {
		w.Header().Set(headerKey, headerValue)
	}
	w.WriteHeader(endpoint.Response.Status)
	w.Write([]byte(endpoint.Response.Body))
}

func notFoundMockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint Not Found: %v", r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}
