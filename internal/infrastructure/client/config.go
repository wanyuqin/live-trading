package client

import "net/http"

type ClientConfig struct {
	BaseURL    string
	HTTPClient *http.Client
}

func DefaultConfig() ClientConfig {
	return ClientConfig{
		HTTPClient: &http.Client{},
	}
}
