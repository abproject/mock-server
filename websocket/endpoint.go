package websocket

import (
	"fmt"
	. "github.com/abproject/mock-server/shared"
	"log"
	"strconv"
	. "strings"
)

type Endpoint struct {
	Id       string
	Url      string
	Interval int
	Repeat   int64
	Order    string
	Messages []string
}

func (endpoint *Endpoint) Init(config EndpointConfig) {
	endpoint.parse(config)
}

func (endpoint *Endpoint) parse(config EndpointConfig) {
	if config.Url == "" {
		log.Fatal(fmt.Sprintf("Websocket Config: Url is required\n%#v", config))
	}

	var order = ToLower(config.Order)
	if order == "" {
		order = "start"
	}
	if order != "start" && order != "end" && order != "random" {
		log.Fatal(fmt.Sprintf("Websocket Config: Order must be either 'start' or 'end' or 'random'\n%#v", config))
	}

	var repeat = ToLower(config.Repeat)
	if repeat == "" {
		repeat = "1"
	}
	if repeat == "infinite" {
		repeat = "9223372036854775807"
	}

	var repeatValue, err = strconv.ParseInt(repeat, 10, 64)


	if err != nil {
		log.Fatal(fmt.Sprintf("Websocket Config: Repeat must be 'infinite' or a number\n%#v", config))
	}

	*endpoint = Endpoint{
		GetRandomId(),
		config.Url,
		config.Interval,
		repeatValue,
		order,
		config.Messages,
	}
}
