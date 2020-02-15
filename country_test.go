package airvisual

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountries(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   []*Countries
		err    error
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
			want: []*Countries{
				{Country: "Andorra"},
				{Country: "Argentina"},
			},
			err: nil,
		},
		{
			name: "countries request failed",
			result: `{
  "status": "call_limit_reached",
  "data": []
}`,
			want: nil,
			err:  fmt.Errorf("unable to list countries: %v", "call_limit_reached"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, err := client.Countries()
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
			if !reflect.DeepEqual(test.err, err) {
				t.Errorf("expected %#v , got %#v", test.err, err)
			}
		})
	}
}
