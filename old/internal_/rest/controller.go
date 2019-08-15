package rest

import (
	. "github.com/abproject/mock-server/internal_/config"
	. "github.com/abproject/mock-server/internal_/shared"
)

type Controller struct {
	Id       string `json:"id"`
	RequestConfig RestRequestConfig `json:"request"`
	ResponseConfig RestResponseConfig `json:"response"`
	Request  Request `json:"-"`
	Response Response `json:"-"`
}

func NewController(config RestControllerConfig) *Controller {
	return &Controller{
		Id:       GetRandomId(),
		RequestConfig: config.Request,
		ResponseConfig: config.Response,
		Request:  *NewRequest(config.Request),
		Response: *NewResponse(config.Response),
	}
}

func (controller *Controller) Patch(config RestControllerConfig) {
	controller.Request.Patch(config.Request)
	controller.Response.Patch(config.Response)
}
