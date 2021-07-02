package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

// APIConfig stores configs parsed from ./config/config.yaml
type APIConfig struct {
	Host     string `yaml:"host"`
	Delegate string `yaml:"delegate"`
}

// GetConfig reads config from the file
func GetConfig(filename string) (config *APIConfig, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var conf APIConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, err
}
