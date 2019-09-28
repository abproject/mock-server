package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/abproject/mock-server/internal/file"

	"github.com/abproject/mock-server/internal/rest"
)

// Context Router Context
type Context struct {
	Logger      *log.Logger
	RestStorage *rest.StorageRest
	FileStorage *file.StorageFile
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
		// REST Endpoints
		if strings.HasPrefix(r.RequestURI, "/_api/rest/endpoints") {
			RouteRestEndpointAPI(*router.context, w, r)
			return
		}
		// REST Global
		if strings.HasPrefix(r.RequestURI, "/_api/rest/global") {
			RouteRestGlobalAPI(*router.context, w, r)
			return
		}
		if strings.HasPrefix(r.RequestURI, "/_api/file") {
			RouteFileAPI(*router.context, w, r)
			return
		}
		// ...
		notFoundHandler(*router.context, w, r)
	} else {
		// Mock
		RouteMock(*router.context, w, r)
	}
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
