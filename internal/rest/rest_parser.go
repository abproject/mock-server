package rest

import (
	"log"
)

// Context Rest Parser Context
type Context struct {
	Logger      *log.Logger
	RestStorage *StorageRest
}

// ParseConfig Parsinf Rest Config
func ParseConfig(c Context, config Config) {
	storage := c.RestStorage
	for _, endpoint := range config.Endpoints {
		(*storage).Add(endpoint)
	}
}
