package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ApiConfig struct {
	Host     string `yaml:"host"`
	Delegate string `yaml:"delegate"`
	Cycle    string `yaml:"cycle"`
}

func GetConfig(filename string) (config *ApiConfig, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	config = new(ApiConfig)
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, err
}
