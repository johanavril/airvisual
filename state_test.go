package airvisual

import (
	"reflect"
	"testing"
)

func TestStates(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *States
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
			want: &States{
				Status: "success",
				Data: []*StatesData{
					{State: "AR"},
					{State: "Alabana"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.States("USA")
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
		})
	}
}
