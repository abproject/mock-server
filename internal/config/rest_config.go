package config

type RestConfig struct {
	Global      RestControllerConfig   `json:"global" yaml:"global"`
	Controllers []RestControllerConfig `json:"endpoints" yaml:"endpoints"`
}

type RestControllerConfig struct {
	Request  RestRequestConfig  `json:"request" yaml:"request"`
	Response RestResponseConfig `json:"response" yaml:"response"`
}

type RestRequestConfig struct {
	Method  string            `json:"method"  yaml:"method"`
	Path    string            `json:"path"    yaml:"path"`
	PathReg string            `json:"pathReg" yaml:"pathReg"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

type RestResponseConfig struct {
	Body    string            `json:"body" yaml:"body"`
	File    string            `json:"file" yaml:"file"`
	Status  int               `json:"status" yaml:"status"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}
