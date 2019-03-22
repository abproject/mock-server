package rest

type RequestConfig struct {
	Method  string            `json:"method"  yaml:"method"`
	Path    string            `json:"path"    yaml:"path"`
	PathReg string            `json:"pathReg" yaml:"pathReg"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}
