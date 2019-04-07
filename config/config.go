package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Rest      RestConfig      `yaml:"rest"`
	Websocket WebsocketConfig `yaml:"websocket"`
}

func ParseConfig(filePath string) *Config {
	config := new(Config)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("File parse err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}
	return config
}