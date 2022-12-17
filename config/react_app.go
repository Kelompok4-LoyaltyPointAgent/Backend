package config

import "os"

type ReactAppConfig struct {
	BaseURL string
}

func LoadReactAppConfig() ReactAppConfig {
	return ReactAppConfig{
		BaseURL: os.Getenv("REACT_APP_BASE_URL"),
	}
}
