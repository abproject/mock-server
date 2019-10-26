package parser

import "github.com/abproject/mock-server/internal/rest/restmodels"

// Config Entry poinf og configuration file
type Config struct {
	Rest restmodels.Config `json:"rest" yaml:"rest"`
}
