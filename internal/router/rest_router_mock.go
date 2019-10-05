package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abproject/mock-server/internal/rest"
)

// RouteMock Rest API
func RouteMock(c Context, w http.ResponseWriter, r *http.Request) {
	endpoint, err := (*c.RestStorage).FindByRequest(r)
	if err != nil {
		notFoundMockHandler(w, r)
		return
	}
	// Rest endpoint found
	for headerKey, headerValue := range endpoint.Response.Headers {
		w.Header().Set(headerKey, headerValue)
	}
	w.WriteHeader(endpoint.Response.Status)
	if endpoint.Response.BodyFile != "" {
		body, err := (*c.FileStorage).GetBody(endpoint.Response.BodyFile)
		if err != nil {
			notFoundMockFileHandler(w, r, endpoint)
			return
		}
		w.Write(body)
	} else {
		w.Write([]byte(endpoint.Response.Body))
	}
}

func notFoundMockHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint Not Found: %v", r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

func notFoundMockFileHandler(w http.ResponseWriter, r *http.Request, endpoint rest.EndpointRestDto) {
	log.Printf("File Not Found: %v", endpoint.Response.BodyFile)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("File '%s' not found", endpoint.Response.BodyFile)))
}
