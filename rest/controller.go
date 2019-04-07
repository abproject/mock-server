package rest

import .  "github.com/abproject/mock-server/shared"

type Controller struct {
	Id       string
	Request  Request
	Response Response
}

func (controller *Controller) Init(config ControllerConfig) {
	controller.Id = GetRandomId()
	controller.Request.Init(config.Request)
	controller.Response.Init(config.Response)
}

func (controller *Controller) Patch(config ControllerConfig) {
	controller.Request.Patch(config.Request)
	controller.Response.Patch(config.Response)
}

