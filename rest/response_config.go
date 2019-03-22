package rest

type ResponseConfig struct {
	Body string `json:"body" yaml:"body"`
	File string `json:"file" yaml:"file"`
	Status int `json:"status" yaml:"status"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}