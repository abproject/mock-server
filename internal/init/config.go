package init

import (
	"github.com/abproject/mock-server/internal/rest"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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