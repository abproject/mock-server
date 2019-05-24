package api

import (
	"encoding/json"
	"errors"
	. "github.com/abproject/mock-server/internal/config"
	"github.com/abproject/mock-server/internal/rest"
	"io/ioutil"
	"log"
	"net/http"
)

var router = Router{}
func init() {
	router.Get("/_api/rest", func (hd HandlerData) {
		json.NewEncoder(hd.w).Encode(rest.GetControllers())
	})

	router.Get("/_api/rest/:id", func (hd HandlerData) {
		id := hd.vars["id"]
		err, controller := rest.GetControllersById(id)
		if err != nil {
			hd.w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(hd.w).Encode(controller)
		}
	})

	router.Post("/_api/rest", func (hd HandlerData) {
		err, config := getConfig(hd)
		if err != nil {
			return
		}
		rest.Add(config)
		hd.w.WriteHeader(http.StatusCreated)
	})

	//router.Put("/_api/rest/:id", func (hd HandlerData) {
	//	id := hd.vars["id"]
	//	body, err := ioutil.ReadAll(hd.r.Body)
	//	if err != nil {
	//		log.Printf("%s: %#v\nInvalid request: %#v", hd.r.RequestURI, hd.r.Body, err)
	//		hd.w.WriteHeader(http.StatusBadRequest)
	//		hd.w.Write([]byte("Invalid request"))
	//		return
	//	}
	//
	//	var config config.RestControllerConfig
	//	err = json.Unmarshal(body, &config)
	//	if err != nil {
	//		log.Printf("%s: %#v\nInvalid request: %#v", hd.r.RequestURI, hd.r.Body, err)
	//		hd.w.WriteHeader(http.StatusBadRequest)
	//		hd.w.Write([]byte("Invalid request"))
	//		return
	//	}
	//	rest.Add(config)
	//	hd.w.WriteHeader(http.StatusCreated)
	//})
}

func Route(w http.ResponseWriter, r *http.Request) {
	router.route(w, r)
}

func getConfig(hd HandlerData) (error, RestControllerConfig)  {
	body, err := ioutil.ReadAll(hd.r.Body)
	if err != nil {
		log.Printf("Can't parse body: %#v", err)
		hd.w.WriteHeader(http.StatusBadRequest)
		hd.w.Write([]byte("Invalid request"))
		return errors.New("Invalid request\n"), RestControllerConfig{}
	}

	var config RestControllerConfig
	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Printf("Can't parse body: %#v\n%#v", hd.r.Body, err)
		hd.w.WriteHeader(http.StatusBadRequest)
		hd.w.Write([]byte("Invalid request"))
		return errors.New("Invalid request\n"), RestControllerConfig{}
	}
	return nil, config
}

