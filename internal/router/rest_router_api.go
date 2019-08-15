package router

import (
	"encoding/json"
	"net/http"

	"github.com/abproject/mock-server/internal/rest"
)

// RouteAPI Rest API
func RouteAPI(c Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == "/_api/rest" {
		switch r.Method {
		case "GET":
			getAllHandler(c, w, r)
			return
		case "POST":
			postHandler(c, w, r)
			return
		}
	}
	notFoundHandler(c, w, r)
}

func getAllHandler(c Context, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*c.RestStorage).GetAll())
}

func postHandler(c Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dto rest.EndpointRestDto
	err := decoder.Decode(&dto)
	if err != nil {
		errorHandler(w, err)
		return
	}
	endpoint := (*c.RestStorage).Add(dto)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(endpoint)
}

func notFoundHandler(c Context, w http.ResponseWriter, r *http.Request) {
	c.Logger.Printf("API Endpoint Not Found\nURI: %s\nMethod: %s", r.RequestURI, r.Method)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
}

func errorHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
