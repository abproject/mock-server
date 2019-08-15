package rest

import (
	"errors"
	. "github.com/abproject/mock-server/internal_/config"
	"net/http"
	"sort"
)

type Rest struct {
	controllers []Controller
	config      RestConfig
}

var storage Rest

func FileRest(config RestConfig) {
	storage.config = config
	for _, controllerConfig := range config.Controllers {
		Add(controllerConfig)
	}
}

func Add(controllerConfig RestControllerConfig) {
	controller := NewController(controllerConfig)
	controller.Patch(storage.config.Global)
	storage.controllers = append(storage.controllers, *controller)
}

func Clear() {
	storage.controllers = nil
	storage.config = RestConfig{}
}


func GetControllers() []Controller {
	tmp := make([]Controller, len(storage.controllers))
	copy(tmp, storage.controllers)
	return tmp
}

func GetControllersById(id string) (error, Controller) {
	for _, controller := range storage.controllers {
		if controller.Id == id {
			return nil, controller
		}
	}
	return errors.New("Not Found\n"), Controller{}
}

func FindController(r *http.Request) (error, Controller) {
	var filtered []Controller
	for _, controller := range storage.controllers {
		if controller.Request.IsEqual(r) {
			filtered = append(filtered, controller)
		}
	}
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Request.CompareTo(&filtered[j].Request)
	})

	if len(filtered) == 0 {
		return errors.New("NO ENDPOINT FOUND"), Controller{}
	}
	return nil, filtered[0]
}
