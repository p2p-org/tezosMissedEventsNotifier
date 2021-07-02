package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ApiConfig struct {
	Host     string `yaml:"host"`
	Delegate string `yaml:"delegate"`
}

func GetConfig(filename string) (config *ApiConfig, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf ApiConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, err
}
