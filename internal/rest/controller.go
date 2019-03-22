package rest

import (
	"encoding/hex"
	"math/rand"
)

type Controller struct {
	Id       string
	Request  Request
	Response Response
}

func (controller *Controller) Init(config ControllerConfig) {
	controller.Id = getRandomId()
	controller.Request.Init(config.Request)
	controller.Response.Init(config.Response)
}

func (controller *Controller) Patch(config ControllerConfig) {
	controller.Request.Patch(config.Request)
	controller.Response.Patch(config.Response)
}

func getRandomId() string {
	u := make([]byte, 16)
	rand.Read(u)
	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F
	return hex.EncodeToString(u)
}