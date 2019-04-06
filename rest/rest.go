package rest

import (
	"errors"
	"net/http"
	"sort"
)

type Rest struct {
	controllers []Controller
	config      RestConfig
}

func (rest *Rest) Init(config RestConfig) {
	var controllers = make([]Controller, len(config.Controllers))

	for index, controllerConfig := range config.Controllers {
		var controller Controller
		controller.Init(controllerConfig)
		controller.Patch(config.Global)
		controllers[index] = controller
	}
	*rest = Rest{
		config: config,
		controllers: controllers,
	}
}

func (rest *Rest) GetControllers() []Controller {
	tmp := make([]Controller, len(rest.controllers))
	copy(tmp, rest.controllers)
	return tmp
}

func (rest *Rest) FindController(r *http.Request) (error, Controller) {
	var filtered []Controller
	for _, controller := range rest.controllers {
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

