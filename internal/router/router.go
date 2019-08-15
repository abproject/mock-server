package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/abproject/mock-server/internal/rest"
)

// Context Router Context
type Context struct {
	Logger      *log.Logger
	RestStorage *rest.StorageRest
}

// Router Router
type Router struct {
	context *Context
}

// IRouter interface
type IRouter interface {
	Route(w http.ResponseWriter, r *http.Request)
}

// New Create new Router with Context
func New(context Context) IRouter {
	return &Router{
		context: &context,
	}
}

// Route Route rest endpoints
func (router *Router) Route(w http.ResponseWriter, r *http.Request) {
	// log.Printf("RequestURI: %+v\n", r.RequestURI)
	// log.Printf("Method: %+v\n", r.Method)
	if strings.HasPrefix(r.RequestURI, "/_api") {
		// API
		// REST
		if strings.HasPrefix(r.RequestURI, "/_api/rest") {
			RouteAPI(*router.context, w, r)
		}
		// ...
	} else {
		// Mock
		RouteMock(*router.context, w, r)
	}
}
