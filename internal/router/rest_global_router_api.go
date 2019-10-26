package router

import (
	"encoding/json"
	"net/http"

	"github.com/abproject/mock-server/internal/rest/restmodels"
)

var restGlobalURL = "/_api/rest/global"

// RouteRestGlobalAPI Rest API
func RouteRestGlobalAPI(c Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.RequestURI == restGlobalURL {
		switch r.Method {
		case "GET":
			getRestGlobalHandler(c, w, r)
			return
		case "POST":
			postRestGlobalHandler(c, w, r)
			return
		case "DELETE":
			deleteRestGlobalHandler(c, w, r)
			return
		}
	}
	notFoundHandler(c, w, r)
}

func getRestGlobalHandler(c Context, w http.ResponseWriter, r *http.Request) {
	global := (*c.RestStorage).GetGlobal()
	json.NewEncoder(w).Encode(global)
}

func postRestGlobalHandler(c Context, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dto restmodels.EndpointRestDto
	err := decoder.Decode(&dto)
	if err != nil {
		errorHandler(w, err)
		return
	}
	global := (*c.RestStorage).AddGlobal(dto)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(global)
}

func deleteRestGlobalHandler(c Context, w http.ResponseWriter, r *http.Request) {
	(*c.RestStorage).DeleteGlobal()
	w.WriteHeader(http.StatusNoContent)
}
