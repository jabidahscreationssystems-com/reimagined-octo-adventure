package config

import (
	"fmt"
	"os"
)

// Config holds the configuration for Agent Cortana
type Config struct {
	Port            string
	GitHubAppID     string
	GitHubAppSecret string
	WebhookSecret   string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appID := os.Getenv("GITHUB_APP_ID")
	if appID == "" {
		return nil, fmt.Errorf("GITHUB_APP_ID environment variable is required")
	}

	appSecret := os.Getenv("GITHUB_APP_SECRET")
	webhookSecret := os.Getenv("WEBHOOK_SECRET")

	return &Config{
		Port:            port,
		GitHubAppID:     appID,
		GitHubAppSecret: appSecret,
		WebhookSecret:   webhookSecret,
	}, nil
}
