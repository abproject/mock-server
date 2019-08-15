package router

import (
	"github.com/abproject/mock-server/internal_/api"
	"github.com/abproject/mock-server/internal_/rest"
	"github.com/abproject/mock-server/internal_/websocket"
	"log"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.RequestURI, "/_api") {
		// API
		api.Route(w, r)
		return
	}
	err, endpoint := websocket.FindEndpoint(r)
	if err == nil {
		// Websocket found
		websocket.Subscribe(w, r, endpoint)
		return
	}
	err, controller := rest.FindController(r)
	if err == nil {
		// Rest endpoint found
		log.Printf("%s: %#v", r.RequestURI, controller)
		for headerKey, headerValue := range controller.Response.Headers {
			w.Header().Set(headerKey, headerValue)
		}
		w.WriteHeader(controller.Response.Status)
		if controller.Response.File != nil {
			w.Write(controller.Response.File)
		} else {
			w.Write([]byte(controller.Response.Body))
		}
		return
	}
	// No endpoint found
	log.Printf("%s: %v", r.RequestURI, err)
	w.WriteHeader(http.StatusNotFound)
}
