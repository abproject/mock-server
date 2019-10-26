package parser

import (
	"io/ioutil"

	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/rest"
	"gopkg.in/yaml.v2"
)

// Parser Parser
type Parser struct {
	context *models.AppContext
}

// IParser interface
type IParser interface {
	Parse(file string)
}

// New Create new Router with Context
func New(context models.AppContext) IParser {
	return &Parser{
		context: &context,
	}
}

// Parse Parse config file
func (parser *Parser) Parse(filePath string) {
	config := parser.makeConfig(filePath)

	restContext := models.AppContext{
		Logger:      parser.context.Logger,
		RestStorage: parser.context.RestStorage,
		FileStorage: parser.context.FileStorage,
		Path:        parser.context.Path,
	}
	rest.ParseConfig(restContext, config.Rest)
}

func (parser *Parser) makeConfig(filePath string) *Config {
	logger := parser.context.Logger
	config := new(Config)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Fatalf("Couldn't read file %s\n%+v", filePath, err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		logger.Fatalf("Couldn't unmarshal file %s:\n%+v", filePath, err)
	}
	logger.Printf("Parsed file: %s", filePath)
	return config
}
