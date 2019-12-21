package airvisual

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func mockClientServer(result string) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(result))
	}))

	client := &Client{
		APIKey:       "API Key",
		baseEndpoint: server.URL,
		client:       server.Client(),
	}

	return client, server
}

func TestEndpoint(t *testing.T) {
	client, server := mockClientServer("")

	defer server.Close()

	tests := []struct {
		name string
		api  string
		v    url.Values
		want string
	}{
		{
			name: "empty param endpoint",
			api:  "empty",
			want: server.URL + "empty?",
		},
		{
			name: "single param endpoint",
			api:  "single",
			v: url.Values{
				"a": {"1"},
			},
			want: server.URL + "single?a=1",
		},
		{
			name: "multi param endpoint",
			api:  "multi",
			v: url.Values{
				"a": {"1"},
				"b": {"2"},
				"c": {"3"},
			},
			want: server.URL + "multi?a=1&b=2&c=3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := client.endpoint(test.api, test.v)
			want := test.want

			if want != got {
				t.Errorf("expected %s , got %s", want, got)
			}
		})
	}
}
