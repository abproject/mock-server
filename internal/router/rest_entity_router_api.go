package router

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/abproject/mock-server/internal/rest/restentity"
)

var restEntityURL = "/_api/rest/entities"

// RouteRestEntityAPI Rest API
func RouteRestEntityAPI(c Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == restEntityURL {
		switch r.Method {
		case "GET":
			getAllRestEntityHandlers(c, w, r)
			return
		case "POST":
			postRestEntityHandler(c, w, r)
			return
		case "DELETE":
			deleteAllRestEntityHandlers(c, w, r)
			return
		}
	} else if strings.HasPrefix(r.RequestURI, restEntityURL) {
		reg, _ := regexp.Compile(restEntityURL + "/([a-zA-Z0-9]+)")
		groups := reg.FindStringSubmatch(r.RequestURI)
		if len(groups) == 2 {
			id := groups[1]
			switch r.Method {
			case "GET":
				getRestEntityHandler(c, w, r, id)
				return
			case "PUT":
				putRestEntityHandler(c, w, r, id)
				return
			case "DELETE":
				deleteRestEntityHandler(c, w, r, id)
				return
			}
		}
	}
	notFoundHandler(c, w, r)
}

func getAllRestEntityHandlers(c Context, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode((*c.RestStorage).GetAllEntities())
}

func getRestEntityHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	entity, err := (*c.RestStorage).GetEntity(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func postRestEntityHandler(c Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dto restentity.EntityRestDto
	err := decoder.Decode(&dto)
	if err != nil {
		errorHandler(w, err)
		return
	}
	entity := (*c.RestStorage).AddEntity(dto)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity)
}

func putRestEntityHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	decoder := json.NewDecoder(r.Body)
	var dto restentity.EntityRestDto
	err := decoder.Decode(&dto)
	if err != nil {
		errorHandler(w, err)
		return
	}
	entity, err := (*c.RestStorage).PutEntity(id, dto)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity)
}

func deleteRestEntityHandler(c Context, w http.ResponseWriter, r *http.Request, id string) {
	err := (*c.RestStorage).DeleteEntity(id)
	if err != nil {
		notFoundHandler(c, w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteAllRestEntityHandlers(c Context, w http.ResponseWriter, r *http.Request) {
	(*c.RestStorage).DeleteAllEntities()
	w.WriteHeader(http.StatusNoContent)
}
