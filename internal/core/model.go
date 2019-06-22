package core

import . "github.com/abproject/mock-server/internal/rest"

type Config struct {
	Rest ConfigRestDto `json:"rest" yaml:"rest"`
	//Websocket WebsocketConfig `json:"websocket" yaml:"websocket"`
}
