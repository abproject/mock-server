package core

import (
	"log"
	"net/http"
	"strings"
)

type Router interface {

}

func CoreRouter(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.RequestURI, "/_api") {
		//_api/rest
		//_api/entity
		//_api/websocket
		// API
		log.Printf("API call %s", r.RequestURI)
		return
	}
	//not _api
	// websocket
}
