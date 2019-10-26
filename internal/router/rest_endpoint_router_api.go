package router

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/abproject/mock-server/internal/rest/restmodels"
)

var restEndpointURL = "/_api/rest/endpoints"

// RouteRestEndpointAPI Rest API
func RouteRestEndpointAPI(c Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == restEndpointURL {
		switch r.Method {
		case "GET":
			getAllRestEndpointHandlers(c, w, r)
			return
		case "POST":
			postRestEndpointHandler(c, w, r)
			return
		case "DELETE":
			deleteAllRestEndpointHandlers(c, w, r)
			return
		}
	} else if strings.HasPrefix(r.RequestURI, restEndpointURL) {
		reg, _ := regexp.Compile(restEndpointURL + "/([a-zA-Z0-9]+)")
		groups := reg.FindStringSubmatch(r.RequestURI)
		if len(groups) == 2 {
			id := groups[1]
			switch r.Method {
			case "GET":
				getRestEndpointHandler(c, w, r, id)
				return
			case "PUT":
				putRestEndpointHandler(c, w, r, id)
				return
			case "DELETE":
				deleteRestEndpointHandler(c, w, r, id)
				return
			}
		}
	}
	notFoundHandler(c, w, r)
}

func getAllRestEndpointHandlers(c Context, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*c.RestStorage).GetAll())
}

func getRestEndpointHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	endpoint, err := (*c.RestStorage).Get(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	json.NewEncoder(w).Encode(endpoint)
}

func postRestEndpointHandler(c Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dto restmodels.EndpointRestDto
	err := decoder.Decode(&dto)
	if err != nil {
		errorHandler(w, err)
		return
	}
	endpoint := (*c.RestStorage).Add(dto)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(endpoint)
}

func putRestEndpointHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	decoder := json.NewDecoder(r.Body)
	var dto restmodels.EndpointRestDto
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

func deleteRestEndpointHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	err := (*c.RestStorage).Delete(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteAllRestEndpointHandlers(c Context, w http.ResponseWriter, r *http.Request) {
	(*c.RestStorage).DeleteAll()
	w.WriteHeader(http.StatusNoContent)
}
