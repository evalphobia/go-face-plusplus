package config

import (
	"os"
	"time"

	"github.com/evalphobia/go-face-plusplus/client"
)

const (
	defaultEnvAPIKey    = "FACEPP_API_KEY"
	defaultEnvAPISecret = "FACEPP_API_SECRET"
)

var (
	envAPIKey    string
	envAPISecret string
)

func init() {
	envAPIKey = os.Getenv(defaultEnvAPIKey)
	envAPISecret = os.Getenv(defaultEnvAPISecret)
}

// Config contains parameters for Face++.
type Config struct {
	APIKey    string
	APISecret string

	Debug   bool
	Timeout time.Duration
}

// Client generates client.Client from Config data.
func (c Config) Client() (*client.Client, error) {
	cli := client.New()
	cli.SetOption(client.Option{
		Debug:   c.Debug,
		Timeout: c.Timeout,
	})
	cli.SetAPIKeys(c.getAPIKeys())
	return cli, nil
}

// getAPIKeys returns API Key and API Secret from Config data or Environment variables.
func (c Config) getAPIKeys() (apiKey, apiSecret string) {
	apiKey = c.APIKey
	if apiKey == "" {
		apiKey = envAPIKey
	}

	apiSecret = c.APISecret
	if apiSecret == "" {
		apiSecret = envAPISecret
	}
	return
}
