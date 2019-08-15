package core

import (
	"fmt"
	"github.com/abproject/mock-server/internal/rest"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
var parserModuleName = "Parser"

type Config struct {
	Rest rest.ConfigRestDto `json:"rest" yaml:"rest"`
}

func ParseFile(filePath string) {
	logger := GetLogger()
	config := new(Config)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error(parserModuleName, err)
		return
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		logger.Error(parserModuleName, err)
		return
	}
	logger.Info(parserModuleName, fmt.Sprintf("File parsed %s", filePath))
}