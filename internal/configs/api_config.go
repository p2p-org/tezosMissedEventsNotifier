package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ApiConfig struct {
	host string `yaml:"host"`
}

func GetApiLink(filename string) (str string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var cfg ApiConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return "", err
	}
	return cfg.host, err
}
