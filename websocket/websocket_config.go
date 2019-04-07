package websocket

type WebsocketConfig struct {
	Endpoints []EndpointConfig `yaml:"endpoints"`
}

func (websocketConfig *WebsocketConfig) Init() Websocket {
	var websocket Websocket
	websocket.Init(*websocketConfig)
	return websocket
}
