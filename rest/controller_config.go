package rest

type ControllerConfig struct {
	Request  RequestConfig  `json:"request" yaml:"request"`
	Response ResponseConfig `json:"response" yaml:"response"`
}
