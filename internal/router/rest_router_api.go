package router

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/abproject/mock-server/internal/rest"
)

var restURL = "/_api/rest/endpoint"

// RouteRestAPI Rest API
func RouteRestAPI(c Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == restURL {
		switch r.Method {
		case "GET":
			getAllRestHandlers(c, w, r)
			return
		case "POST":
			postRestHandler(c, w, r)
			return
		case "DELETE":
			deleteAllRestHandlers(c, w, r)
			return
		}
	} else if strings.HasPrefix(r.RequestURI, restURL) {
		reg, _ := regexp.Compile(restURL + "/([a-zA-Z0-9]+)")
		groups := reg.FindStringSubmatch(r.RequestURI)
		if len(groups) == 2 {
			id := groups[1]
			switch r.Method {
			case "GET":
				getRestHandler(c, w, r, id)
				return
			case "PUT":
				putRestHandler(c, w, r, id)
				return
			case "DELETE":
				deleteRestHandler(c, w, r, id)
				return
			}
		}
	}
	notFoundHandler(c, w, r)
}

func getAllRestHandlers(c Context, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*c.RestStorage).GetAll())
}

func getRestHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	endpoint, err := (*c.RestStorage).Get(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	json.NewEncoder(w).Encode(endpoint)
}

func postRestHandler(c Context, w http.ResponseWriter, r *http.Request) {
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

func putRestHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	decoder := json.NewDecoder(r.Body)
	var dto rest.EndpointRestDto
	err := decoder.Decode(&dto)
	if err != nil {
		errorHandler(w, err)
		return
	}
	endpoint, err := (*c.RestStorage).Put(id, dto)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(endpoint)
}

func deleteRestHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	err := (*c.RestStorage).Delete(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteAllRestHandlers(c Context, w http.ResponseWriter, r *http.Request) {
	(*c.RestStorage).DeleteAll()
	w.WriteHeader(http.StatusNoContent)
}
