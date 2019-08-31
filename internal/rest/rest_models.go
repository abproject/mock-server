package rest

//Config Entry point for rest config
type Config struct {
	Endpoints []EndpointRestDto `json:"endpoints" yaml:"endpoints" `
}

// EndpointRestDto Rest Endpoint Data Transfer Object
type EndpointRestDto struct {
	ID       string          `json:"id" yaml:"id"`
	Request  RequestRestDto  `json:"request" yaml:"request"`
	Response ResponseRestDto `json:"response" yaml:"response"`
}

// RequestRestDto Rest Request Data Transfer Object
type RequestRestDto struct {
	Method  string            `json:"method"  yaml:"method"`
	Path    string            `json:"path"    yaml:"path"`
	PathReg string            `json:"pathReg" yaml:"pathReg"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

// ResponseRestDto Rest Response Data Transfer Object
type ResponseRestDto struct {
	Body     string            `json:"body" yaml:"body"`
	BodyFile string            `json:"bodyFile" yaml:"bodyFile"`
	Status   int               `json:"status" yaml:"status"`
	Headers  map[string]string `json:"headers" yaml:"headers"`
}
