package parser

import (
	"io/ioutil"
	"log"

	"github.com/abproject/mock-server/internal/rest"
	"gopkg.in/yaml.v2"
)

// Context Parser Context
type Context struct {
	Logger      *log.Logger
	RestStorage *rest.StorageRest
}

// Parser Parser
type Parser struct {
	context *Context
}

// IParser interface
type IParser interface {
	Parse(file string)
}

// New Create new Router with Context
func New(context Context) IParser {
	return &Parser{
		context: &context,
	}
}

// Parse Parse config file
func (parser *Parser) Parse(filePath string) {
	config := parser.makeConfig(filePath)

	restContext := rest.Context{
		Logger:      parser.context.Logger,
		RestStorage: parser.context.RestStorage,
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
