package airvisual

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStates(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   []*States
		err    error
	}{
		{
			name: "states request success",
			result: `{
  "status": "success",
  "data": [
    {
      "state": "AR"
    },
    {
      "state": "Alabana"
    }
  ]
}`,
			want: []*States{
				{State: "AR"},
				{State: "Alabana"},
			},
			err: nil,
		},
		{
			name: "states request failed",
			result: `{
  "status": "call_limit_reached",
  "data": []
}`,
			want: nil,
			err:  fmt.Errorf("unable to list states: %v", "call_limit_reached"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, err := client.States("USA")
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
