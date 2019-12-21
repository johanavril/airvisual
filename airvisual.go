// Package airvisual provides wrapper function for AirVisual's API
package airvisual

import (
	"net/http"
)

// Client is a client to work with airvisual API
type Client struct {
	client       *http.Client
	baseEndpoint string

	APIKey string
}

// New takes API Key and return a client that will utilize that key
func New(apiKey string, opts ...Option) *Client {
	client := Client{
		client:       http.DefaultClient,
		baseEndpoint: baseEndpoint,
		APIKey:       apiKey,
	}

	for _, opt := range opts {
		opt(&client)
	}

	return &client
}

// Option is an option to configure airvisual client
type Option func(*Client)

// WithHTTPClient set http client on airvisual client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}
