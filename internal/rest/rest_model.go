package rest

type RequestRestDto struct {
	Method  string            `json:"method"  yaml:"method"`
	Path    string            `json:"path"    yaml:"path"`
	PathReg string            `json:"pathReg" yaml:"pathReg"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

type ResponseRestDto struct {
	Body    string            `json:"body" yaml:"body"`
	File    string            `json:"file" yaml:"file"`
	Status  int               `json:"status" yaml:"status"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

type EndpointRestDto struct {
	Request  RequestRestDto  `json:"request" yaml:"request"`
	Response ResponseRestDto `json:"response" yaml:"response"`
}

type EndpointOutRestDto struct {
	Id string `json:"id" yaml:"id"`
	EndpointRestDto
}

type ConfigRestDto struct {
	Global    EndpointRestDto   `json:"global" yaml:"global"`
	Endpoints []EndpointRestDto `json:"endpoints" yaml:"endpoints"`
}

type requestRestParsed struct {
	method    string
	path      string
	isPathReg bool
	headers   map[string][]string
}

type responseRestParsed struct {
	body    string
	file    []byte
	status  int
	headers map[string]string
}

type endpointRestParsed struct {
	request  requestRestParsed
	response responseRestParsed
}

type restEntry struct {
	config   EndpointRestDto
	endpoint endpointRestParsed
}

