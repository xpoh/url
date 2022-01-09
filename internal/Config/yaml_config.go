package Config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type yamlConfig struct{}

func (y yamlConfig) load() Cfg {
	cfg := Cfg{}
	yamlFile, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		log.Panicln(err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	log.Println("Read yaml file:%v", yamlFile)
	if err != nil {
		log.Panicln(err)
	}
	return cfg
}
