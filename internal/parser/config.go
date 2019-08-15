package parser

import "github.com/abproject/mock-server/internal/rest"

// Config Entry poinf og configuration file
type Config struct {
	Rest rest.Config `json:"rest" yaml:"rest"`
}
