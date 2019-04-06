package init

import (
	. "github.com/abproject/mock-server/rest"
	"log"
	"net/http"
)

type Router struct {
	Rest Rest
}

func (router *Router) Init(config Config) {
	router.Rest = config.RestConfig.Init()
}

func (router *Router) Request(w http.ResponseWriter, r *http.Request) {
	var err, controller = router.Rest.FindController(r)
	if err != nil {
		log.Printf("%s: %v", r.RequestURI, err)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write(nil)
	} else {
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
