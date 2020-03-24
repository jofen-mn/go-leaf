package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configFile []byte
var Configure Config

type Config struct {
	Application Application `yaml:"application"`
	Mysql       Mysql       `yaml:"mysql"`
}

type Application struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type Mysql struct {
	Url          string `yaml:"url"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}

func ConfInit() {
	configPath := "conf/config.yml"

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicf("Config file not found: %s\n", configPath)
		return
	}

	err = yaml.Unmarshal(configFile, &Configure)
	if err != nil {
		log.Panicf("Config file type is error")
		return
	}
}
