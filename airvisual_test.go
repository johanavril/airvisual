package airvisual

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	tests := []struct {
		name string
		got  *Client
		want *Client
	}{
		{
			name: "default client",
			got:  New("API Key"),
			want: &Client{
				client: http.DefaultClient,
				APIKey: "API Key",
			},
		},
		{
			name: "client with option",
			got: New(
				"API Key",
				WithHTTPClient(&http.Client{Timeout: 5 * time.Second}),
			),
			want: &Client{
				client: &http.Client{Timeout: 5 * time.Second},
				APIKey: "API Key",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.got
			want := test.want

			assert.Equal(t, want.client, got.client)
			assert.Equal(t, want.APIKey, got.APIKey)
		})
	}
}
