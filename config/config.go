package config

import (
	"github.com/daffarg/distributed-cascading-cb/util"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Config struct {
	AlternativeEndpoints map[string]AlternativeEndpoint `yaml:"alternativeEndpoints"`
}

type config struct {
	AlternativeEndpoints []AlternativeEndpoint `yaml:"alternativeEndpoints"`
}

type AlternativeEndpoint struct {
	Endpoint            string `yaml:"endpoint"`
	Method              string `yaml:"method"`
	AlternativeEndpoint string `yaml:"alternativeEndpoint"`
	AlternativeMethod   string `yaml:"alternativeMethod"`
}

func NewConfig() *Config {
	return &Config{
		AlternativeEndpoints: make(map[string]AlternativeEndpoint),
	}
}

func (c *Config) Read(configPath string) error {
	tmpConfig := &config{}

	buf, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, tmpConfig)
	if err != nil {
		return err
	}

	for _, ep := range tmpConfig.AlternativeEndpoints {
		parsedUrl, err := util.GetGeneralURLFormat(ep.Endpoint)
		if err != nil {
			return err
		}
		ep.Method = strings.ToUpper(ep.Method)
		ep.Endpoint = util.FormEndpointName(strings.ToLower(parsedUrl), ep.Method)

		parsedUrl, err = util.GetGeneralURLFormat(ep.Endpoint)
		if err != nil {
			return err
		}
		ep.AlternativeMethod = strings.ToUpper(ep.AlternativeMethod)
		ep.AlternativeEndpoint = util.FormEndpointName(strings.ToLower(parsedUrl), ep.AlternativeMethod)

		c.AlternativeEndpoints[ep.Endpoint] = ep
	}

	return err
}
