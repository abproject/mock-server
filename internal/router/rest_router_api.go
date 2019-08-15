package router

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/abproject/mock-server/internal/rest"
)

var url = "/_api/rest/endpoint"

// RouteAPI Rest API
func RouteAPI(c Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == url {
		switch r.Method {
		case "GET":
			getAllHandler(c, w, r)
			return
		case "POST":
			postHandler(c, w, r)
			return
		case "DELETE":
			deleteAllHandler(c, w, r)
			return
		}
	} else if strings.HasPrefix(r.RequestURI, url) {
		reg, _ := regexp.Compile(url + "/([a-zA-Z0-9]+)")
		groups := reg.FindStringSubmatch(r.RequestURI)
		if len(groups) == 2 {
			id := groups[1]
			switch r.Method {
			case "GET":
				getHandler(c, w, r, id)
				return
			case "PUT":
				putHandler(c, w, r, id)
				return
			case "DELETE":
				deleteHandler(c, w, r, id)
				return
			}
		}
	}
	notFoundHandler(c, w, r)
}

func getAllHandler(c Context, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*c.RestStorage).GetAll())
}

func getHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	endpoint, err := (*c.RestStorage).Get(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	json.NewEncoder(w).Encode(endpoint)
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

func putHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
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

func deleteHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	err := (*c.RestStorage).Delete(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteAllHandler(c Context, w http.ResponseWriter, r *http.Request) {
	(*c.RestStorage).DeleteAll()
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
