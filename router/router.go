package router

import (
	. "github.com/abproject/mock-server/config"
	. "github.com/abproject/mock-server/rest"
	. "github.com/abproject/mock-server/websocket"
	"log"
	"net/http"
)

type Router struct {
	Rest Rest
	Websocket Websocket
}

func NewRouter(config Config) *Router {
	return &Router{
		Rest: *NewRest(config.Rest),
		Websocket: *NewWebsocket(config.Websocket),
	}
}

func (router *Router) Request(w http.ResponseWriter, r *http.Request) {
	var err, endpoint = router.Websocket.FindEndpoint(r)
	if err == nil {
		// Websocket found
		router.Websocket.Subscribe(w, r, endpoint)
	} else {
		var err, controller = router.Rest.FindController(r)
		if err != nil {
			// No endpoint found
			log.Printf("%s: %v", r.RequestURI, err)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			w.Write(nil)
		} else {
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
		}
	}
}
