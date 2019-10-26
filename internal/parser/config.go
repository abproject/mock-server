package parser

import "github.com/abproject/mock-server/internal/models"

// Config Entry poinf og configuration file
type Config struct {
	Rest models.Config `json:"rest" yaml:"rest"`
}
