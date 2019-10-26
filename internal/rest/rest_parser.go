package rest

import (
	"io/ioutil"
	"path/filepath"

	"github.com/abproject/mock-server/internal/models"
)

// ParseConfig Parsing Rest models.Config
func ParseConfig(c models.AppContext, config models.Config) {
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

	for _, entity := range config.Entities {
		dataFile := entity.Data
		if dataFile != "" && !(*fileStorage).IsExist(dataFile) {
			path := filepath.Join(c.Path, dataFile)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				c.Logger.Printf("File not found: %s", path)
			} else {
				(*fileStorage).Put(dataFile, dataFile, data)
			}
		}
		newItemFile := entity.NewEntity
		if newItemFile != "" && !(*fileStorage).IsExist(newItemFile) {
			path := filepath.Join(c.Path, newItemFile)
			data, err := ioutil.ReadFile(path)
			if err != nil {
				c.Logger.Printf("File not found: %s", path)
			} else {
				(*fileStorage).Put(newItemFile, newItemFile, data)
			}
		}
		(*restStorage).AddEntity(entity)
	}
}
