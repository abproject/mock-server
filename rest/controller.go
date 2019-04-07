package rest

import (
	. "github.com/abproject/mock-server/config"
	. "github.com/abproject/mock-server/shared"
)

type Controller struct {
	Id       string
	Request  Request
	Response Response
}

func NewController(config RestControllerConfig) *Controller {
	return &Controller {
		Id: GetRandomId(),
		Request: *NewRequest(config.Request),
		Response: *NewResponse(config.Response),
	}
}

func (controller *Controller) Patch(config RestControllerConfig) {
	controller.Request.Patch(config.Request)
	controller.Response.Patch(config.Response)
}

