package config

import (
	"fmt"
	"github.com/google/go-github/v41/github"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// GitHub API client
	GitHubClient        *github.Client
	GitHubWebhookSecret string `mapstructure:"github_webhook_secret"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("go_github_app")
	v.AutomaticEnv()

	v.SetDefault("server_port", 8080)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
