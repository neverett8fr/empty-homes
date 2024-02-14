package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	route = "config/config.yaml"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type HTML struct {
	StyleRoute string `yaml:"style-route"`
	Style      string
}

type DB struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

type Config struct {
	Service Service `yaml:"api-config"`
	DB      DB      `yaml:"db-config"`
	HTML    HTML    `yaml:"html-config"`
}

func Initialise() (*Config, error) {

	conf := Config{}

	yamlFile, err := ioutil.ReadFile(route)
	if err != nil {
		return &Config{}, fmt.Errorf("issue finding config yaml, err %v", err)
	}
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return &Config{}, fmt.Errorf("issue unmarshalling config yaml, err %v", err)
	}

	styleFile, err := ioutil.ReadFile(conf.HTML.StyleRoute)
	if err != nil {
		return &conf, fmt.Errorf("issue finding style file, err %v", err)
	}
	conf.HTML.Style = string(styleFile)

	return &conf, nil
}
