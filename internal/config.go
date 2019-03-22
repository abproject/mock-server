package internal

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"mock-server/internal/rest"
)

type Config struct {
	RestConfig rest.RestConfig `yaml:"rest"`
}

func (config *Config) Parse(filePath string) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}