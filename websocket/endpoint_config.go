package websocket

type EndpointConfig struct {
	Url      string   `json:"url" yaml:"url"`
	Interval int      `json:"interval" yaml:"interval"`
	Repeat   string   `json:"repeat" yaml:"repeat"`
	Order    string   `json:"order" yaml:"order"`
	Messages []string `json:"messages" yaml:"messages"`
}
