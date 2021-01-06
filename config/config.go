package config

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type PlatformConfig struct {
	HttpClientConfig struct {
		Timeout      time.Duration `yaml: timeout`
		Proxy        string        `yaml: "proxy"`
		MaxIdleConns int           `yaml: "max_idle_conns"`
	}

	Binance struct {
		ApiKey    string `yaml:"api_key"`
		SecretKey string `yaml:"secret_key"`
	}
}

func ParseConfigFile(filename string) *PlatformConfig {
	config := &PlatformConfig{}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}
