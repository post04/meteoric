package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct ...
type Config struct {
	GeneralConfig struct {
		Token   string `yaml:"token"`
		Bot     bool   `yaml:"bot"`
		Logging struct {
			WebhookURL string `yaml:"webhook_url"`
			FileName   string `yaml:"file_name"`
		} `yaml:"logging"`
		EnvConfig struct {
			UseEnvFile     bool   `yaml:"use_env_file"`
			TokenValueName string `yaml:"token_value_name"`
		} `yaml:"envConfig"`
	} `yaml:"generalConfig"`
}

// GetConfig loads config.yaml into config struct.
func GetConfig() *Config {
	var config *Config

	if yamlf, err := ioutil.ReadFile("config.yaml"); err == nil {
		err := yaml.Unmarshal(yamlf, &config)
		if err != nil {
			return nil
		}

		return config
	}

	return nil
}
