package init

import (
	"github.com/abproject/mock-server/rest"
	"github.com/abproject/mock-server/websocket"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	RestConfig rest.RestConfig `yaml:"rest"`
	WebsocketConfig websocket.WebsocketConfig `yaml:"websocket"`
}

func (config *Config) Parse(filePath string) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("File parse err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}
}