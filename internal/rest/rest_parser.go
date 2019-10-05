package rest

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/abproject/mock-server/internal/file"
)

// Context Rest Parser Context
type Context struct {
	Logger      *log.Logger
	RestStorage *StorageRest
	FileStorage *file.StorageFile
	Path        string
}

// ParseConfig Parsing Rest Config
func ParseConfig(c Context, config Config) {
	restStorage := c.RestStorage
	fileStorage := c.FileStorage
	(*restStorage).AddGlobal(config.Global)
	for _, endpoint := range config.Endpoints {
		file := endpoint.Response.BodyFile
		if file != "" && !(*fileStorage).IsExist(file) {
			path := filepath.Join(c.Path, file)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				c.Logger.Printf("File not found: %s", path)
			} else {
				(*fileStorage).Put(file, file, data)
			}
		}
		(*restStorage).Add(endpoint)
	}
}
