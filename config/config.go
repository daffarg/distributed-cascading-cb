package config

import (
	"github.com/daffarg/distributed-cascading-cb/util"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Config struct {
	AlternativeEndpoints map[string]AlternativeEndpoint `yaml:"alternativeEndpoints" json:"alternative_endpoints"`
	Exceptions           map[string]Endpoint            `yaml:"exceptions" json:"exceptions"`
}

type config struct {
	AlternativeEndpoints []AlternativeEndpoint `yaml:"alternativeEndpoints" json:"alternative_endpoints"`
	Exceptions           []Endpoint            `yaml:"exceptions" json:"exceptions"`
}

type AlternativeEndpoint struct {
	Endpoint     string     `yaml:"endpoint" json:"endpoint"`
	Method       string     `yaml:"method" json:"method"`
	Alternatives []Endpoint `yaml:"alternatives" json:"alternatives"`
}

type Endpoint struct {
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	Method   string `yaml:"method" json:"method"`
}

func NewConfig() *Config {
	return &Config{
		AlternativeEndpoints: make(map[string]AlternativeEndpoint),
		Exceptions:           make(map[string]Endpoint),
	}
}

func (c *Config) Read(configPath string) error {
	tmpConfig := &config{}

	buf, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
	}

	err = yaml.Unmarshal(buf, tmpConfig)
	if err != nil {
		return err
	}

	for i := range tmpConfig.AlternativeEndpoints {
		tmpConfig.AlternativeEndpoints[i].Method = strings.ToUpper(tmpConfig.AlternativeEndpoints[i].Method)
		parsedUrl, err := util.GetGeneralURLFormat(strings.ToLower(tmpConfig.AlternativeEndpoints[i].Endpoint))
		if err != nil {
			return err
		}
		key := util.FormEndpointName(parsedUrl, tmpConfig.AlternativeEndpoints[i].Method)

		for j := range tmpConfig.AlternativeEndpoints[i].Alternatives {
			tmpConfig.AlternativeEndpoints[i].Alternatives[j].Method = strings.ToUpper(tmpConfig.AlternativeEndpoints[i].Alternatives[j].Method)
			parsedUrl, err = util.GetGeneralURLFormat(strings.ToLower(tmpConfig.AlternativeEndpoints[i].Alternatives[j].Endpoint))
			if err != nil {
				return err
			}
		}

		c.AlternativeEndpoints[key] = tmpConfig.AlternativeEndpoints[i]
	}

	for i := range tmpConfig.Exceptions {
		tmpConfig.Exceptions[i].Method = strings.ToUpper(tmpConfig.Exceptions[i].Method)
		parsedUrl, err := util.GetGeneralURLFormat(strings.ToLower(tmpConfig.Exceptions[i].Endpoint))
		if err != nil {
			return err
		}
		key := util.FormEndpointName(parsedUrl, tmpConfig.Exceptions[i].Method)

		c.Exceptions[key] = tmpConfig.Exceptions[i]
	}

	return err
}
