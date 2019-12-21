package airvisual

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountries(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *Countries
	}{
		{
			name: "countries request success",
			result: `{
  "status": "success",
  "data": [
    {
      "country": "Andorra"
    },
    {
      "country": "Argentina"
    }
  ]
}`,
			want: &Countries{
				Status: "success",
				Data: []*CountriesData{
					{Country: "Andorra"},
					{Country: "Argentina"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.Countries()
			want := test.want

			assert.Equal(t, want, got)
		})
	}
}
