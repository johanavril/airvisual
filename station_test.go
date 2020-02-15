package airvisual

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStations(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   []*Stations
		err    error
	}{
		{
			name: "stations request success",
			result: `{
  "status": "success",
  "data": [
    {
      "location": {
        "type": "Point",
        "coordinates": [
          116.466258,
          39.954352
        ]
      },
      "station": "US Embassy in Beijing"
    },
    {
      "location": {
        "type": "Point",
        "coordinates": [
        116.2148532181,
        40.0078007235
        ]
      },
      "station": "Botanical Garden"
    }
  ]
}`,
			want: []*Stations{
				{
					Location: &Location{
						Type:        "Point",
						Coordinates: []float64{116.466258, 39.954352},
					},
					Station: "US Embassy in Beijing",
				},
				{
					Location: &Location{
						Type:        "Point",
						Coordinates: []float64{116.2148532181, 40.0078007235},
					},
					Station: "Botanical Garden",
				},
			},
			err: nil,
		},
		{
			name: "stations request failed",
			result: `{
  "status": "call_limit_reached",
  "data": []
}`,
			want: nil,
			err:  fmt.Errorf("unable to list stations: %v", "call_limit_reached"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, err := client.Stations("Addison", "New York", "USA")
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

func TestStation(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *Station
		err    error
	}{
		{
			name: "station request success",
			result: `{
  "status": "success",
  "data": {
    "name": "US Embassy in Beijing",
    "city": "Beijing",
    "state": "Beijing",
    "country": "China",
    "location": {
      "type": "Point",
      "coordinates": [
        116.466258,
        39.954352
      ]
    },
    "forecasts": [
      {
        "ts": "2019-08-05T03:00:00.000Z",
        "aqius": 41,
        "aqicn": 14,
        "tp": 25,
        "tp_min": 25,
        "pr": 962,
        "hu": 65,
        "ws": 1,
        "wd": 228,
        "ic": "03n"
      },
      {
        "ts": "2019-08-07T00:00:00.000Z",
        "aqius": 68,
        "aqicn": 29
      }
    ],
    "current": {
      "weather": {
        "ts": "2019-08-01T23:00:00.000Z",
        "tp": 37,
        "pr": 1007,
        "hu": 14,
        "ws": 1,
        "wd": 110,
        "ic": "01d"
      },
      "pollution": {
        "ts": "2019-08-04T19:00:00.000Z",
        "aqius": 70,
        "mainus": "p2",
        "aqicn": 30,
        "maincn": "p2",
        "p2": {
          "conc": 21,
          "aqius": 70,
          "aqicn": 30
        },
        "p1": {
          "conc": 30,
          "aqius": 27,
          "aqicn": 30
        },
        "o3": {
          "conc": 48,
          "aqius": 38,
          "aqicn": 30
        },
        "n2": {
          "conc": 8,
          "aqius": 2,
          "aqicn": 8
        },
        "s2": {
          "conc": 1,
          "aqius": 1,
          "aqicn": 3
        },
        "co": {
          "conc": 0.2,
          "aqius": 2,
          "aqicn": 2
        }
      }
    },
    "history": {
      "weather": [
        {
          "ts": "2019-08-01T23:00:00.000Z",
          "tp": 37,
          "pr": 1007,
          "hu": 14,
          "ws": 1,
          "wd": 110,
          "ic": "01d"
        },
        {
          "ts": "2019-08-01T04:00:00.000Z",
          "tp": 31,
          "pr": 1005,
          "hu": 26,
          "ws": 1,
          "wd": 40,
          "ic": "01n"
        }
      ],
      "pollution": [
        {
          "ts": "2019-08-04T19:00:00.000Z",
          "aqius": 70,
          "mainus": "p2",
          "aqicn": 30,
          "maincn": "p2",
          "p2": {
            "conc": 21,
            "aqius": 70,
            "aqicn": 30
          },
          "p1": {
            "conc": 30,
            "aqius": 27,
            "aqicn": 30
          },
          "o3": {
            "conc": 48,
            "aqius": 38,
            "aqicn": 30
          },
          "n2": {
            "conc": 8,
            "aqius": 2,
            "aqicn": 8
          },
          "s2": {
            "conc": 1,
            "aqius": 1,
            "aqicn": 3
          },
          "co": {
            "conc": 0.2,
            "aqius": 2,
            "aqicn": 2
          }
        },
        {
          "ts": "2019-08-04T18:00:00.000Z",
          "aqius": 57,
          "mainus": "p2",
          "aqicn": 28,
          "maincn": "o3",
          "p2": {
            "conc": 15,
            "aqius": 57,
            "aqicn": 21
          },
          "p1": {
            "conc": 22,
            "aqius": 20,
            "aqicn": 22
          },
          "o3": {
            "conc": 45,
            "aqius": 36,
            "aqicn": 28
          },
          "n2": {
            "conc": 8,
            "aqius": 2,
            "aqicn": 8
          },
          "co": {
            "conc": 0.2,
            "aqius": 2,
            "aqicn": 2
          }
        }
      ]
    }
  }
}`,
			want: &Station{
				Name:    "US Embassy in Beijing",
				City:    "Beijing",
				State:   "Beijing",
				Country: "China",
				Location: &Location{
					Type:        "Point",
					Coordinates: []float64{116.466258, 39.954352},
				},
				Forecasts: []*Forecast{
					{
						TS:    "2019-08-05T03:00:00.000Z",
						AQIUS: 41,
						AQICN: 14,
						TP:    25,
						TPMin: 25,
						PR:    962,
						HU:    65,
						WS:    1,
						WD:    228,
						IC:    "03n",
					},
					{
						TS:    "2019-08-07T00:00:00.000Z",
						AQIUS: 68,
						AQICN: 29,
					},
				},
				Current: &Current{
					Weather: &Weather{
						TS: "2019-08-01T23:00:00.000Z",
						TP: 37,
						PR: 1007,
						HU: 14,
						WS: 1,
						WD: 110,
						IC: "01d",
					},
					Pollution: &Pollution{
						TS:     "2019-08-04T19:00:00.000Z",
						AQIUS:  70,
						MAINUS: "p2",
						AQICN:  30,
						MAINCN: "p2",
						P2: &Unit{
							CONC:  21,
							AQIUS: 70,
							AQICN: 30,
						},
						P1: &Unit{
							CONC:  30,
							AQIUS: 27,
							AQICN: 30,
						},
						O3: &Unit{
							CONC:  48,
							AQIUS: 38,
							AQICN: 30,
						},
						N2: &Unit{
							CONC:  8,
							AQIUS: 2,
							AQICN: 8,
						},
						S2: &Unit{
							CONC:  1,
							AQIUS: 1,
							AQICN: 3,
						},
						CO: &Unit{
							CONC:  0.2,
							AQIUS: 2,
							AQICN: 2,
						},
					},
				},
				History: &History{
					Weather: []*Weather{
						{
							TS: "2019-08-01T23:00:00.000Z",
							TP: 37,
							PR: 1007,
							HU: 14,
							WS: 1,
							WD: 110,
							IC: "01d",
						},
						{
							TS: "2019-08-01T04:00:00.000Z",
							TP: 31,
							PR: 1005,
							HU: 26,
							WS: 1,
							WD: 40,
							IC: "01n",
						},
					},
					Pollution: []*Pollution{
						{
							TS:     "2019-08-04T19:00:00.000Z",
							AQIUS:  70,
							MAINUS: "p2",
							AQICN:  30,
							MAINCN: "p2",
							P2: &Unit{
								CONC:  21,
								AQIUS: 70,
								AQICN: 30,
							},
							P1: &Unit{
								CONC:  30,
								AQIUS: 27,
								AQICN: 30,
							},
							O3: &Unit{
								CONC:  48,
								AQIUS: 38,
								AQICN: 30,
							},
							N2: &Unit{
								CONC:  8,
								AQIUS: 2,
								AQICN: 8,
							},
							S2: &Unit{
								CONC:  1,
								AQIUS: 1,
								AQICN: 3,
							},
							CO: &Unit{
								CONC:  0.2,
								AQIUS: 2,
								AQICN: 2,
							},
						},
						{
							TS:     "2019-08-04T18:00:00.000Z",
							AQIUS:  57,
							MAINUS: "p2",
							AQICN:  28,
							MAINCN: "o3",
							P2: &Unit{
								CONC:  15,
								AQIUS: 57,
								AQICN: 21,
							},
							P1: &Unit{
								CONC:  22,
								AQIUS: 20,
								AQICN: 22,
							},
							O3: &Unit{
								CONC:  45,
								AQIUS: 36,
								AQICN: 28,
							},
							N2: &Unit{
								CONC:  8,
								AQIUS: 2,
								AQICN: 8,
							},
							CO: &Unit{
								CONC:  0.2,
								AQIUS: 2,
								AQICN: 2,
							},
						},
					},
				},
			},
			err: nil,
		},
		{
			name: "station request failed",
			result: `{
  "status": "call_limit_reached",
  "data": null
}`,
			want: nil,
			err:  fmt.Errorf("unable to retrieve station data: %v", "call_limit_reached"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, err := client.Station("US Embassy in Beijing", "Beijing", "Beijing", "China")
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

func TestNearestStationIP(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *Station
		err    error
	}{
		{
			name: "nearest station by IP request success",
			result: `{
  "status": "success",
  "data": {
    "name": "US Embassy in Beijing",
    "city": "Beijing",
    "state": "Beijing",
    "country": "China",
    "location": {
      "type": "Point",
      "coordinates": [
        116.466258,
        39.954352
      ]
    },
    "forecasts": [
      {
        "ts": "2019-08-05T03:00:00.000Z",
        "aqius": 41,
        "aqicn": 14,
        "tp": 25,
        "tp_min": 25,
        "pr": 962,
        "hu": 65,
        "ws": 1,
        "wd": 228,
        "ic": "03n"
      },
      {
        "ts": "2019-08-07T00:00:00.000Z",
        "aqius": 68,
        "aqicn": 29
      }
    ],
    "current": {
      "weather": {
        "ts": "2019-08-01T23:00:00.000Z",
        "tp": 37,
        "pr": 1007,
        "hu": 14,
        "ws": 1,
        "wd": 110,
        "ic": "01d"
      },
      "pollution": {
        "ts": "2019-08-04T19:00:00.000Z",
        "aqius": 70,
        "mainus": "p2",
        "aqicn": 30,
        "maincn": "p2",
        "p2": {
          "conc": 21,
          "aqius": 70,
          "aqicn": 30
        },
        "p1": {
          "conc": 30,
          "aqius": 27,
          "aqicn": 30
        },
        "o3": {
          "conc": 48,
          "aqius": 38,
          "aqicn": 30
        },
        "n2": {
          "conc": 8,
          "aqius": 2,
          "aqicn": 8
        },
        "s2": {
          "conc": 1,
          "aqius": 1,
          "aqicn": 3
        },
        "co": {
          "conc": 0.2,
          "aqius": 2,
          "aqicn": 2
        }
      }
    },
    "history": {
      "weather": [
        {
          "ts": "2019-08-01T23:00:00.000Z",
          "tp": 37,
          "pr": 1007,
          "hu": 14,
          "ws": 1,
          "wd": 110,
          "ic": "01d"
        },
        {
          "ts": "2019-08-01T04:00:00.000Z",
          "tp": 31,
          "pr": 1005,
          "hu": 26,
          "ws": 1,
          "wd": 40,
          "ic": "01n"
        }
      ],
      "pollution": [
        {
          "ts": "2019-08-04T19:00:00.000Z",
          "aqius": 70,
          "mainus": "p2",
          "aqicn": 30,
          "maincn": "p2",
          "p2": {
            "conc": 21,
            "aqius": 70,
            "aqicn": 30
          },
          "p1": {
            "conc": 30,
            "aqius": 27,
            "aqicn": 30
          },
          "o3": {
            "conc": 48,
            "aqius": 38,
            "aqicn": 30
          },
          "n2": {
            "conc": 8,
            "aqius": 2,
            "aqicn": 8
          },
          "s2": {
            "conc": 1,
            "aqius": 1,
            "aqicn": 3
          },
          "co": {
            "conc": 0.2,
            "aqius": 2,
            "aqicn": 2
          }
        },
        {
          "ts": "2019-08-04T18:00:00.000Z",
          "aqius": 57,
          "mainus": "p2",
          "aqicn": 28,
          "maincn": "o3",
          "p2": {
            "conc": 15,
            "aqius": 57,
            "aqicn": 21
          },
          "p1": {
            "conc": 22,
            "aqius": 20,
            "aqicn": 22
          },
          "o3": {
            "conc": 45,
            "aqius": 36,
            "aqicn": 28
          },
          "n2": {
            "conc": 8,
            "aqius": 2,
            "aqicn": 8
          },
          "co": {
            "conc": 0.2,
            "aqius": 2,
            "aqicn": 2
          }
        }
      ]
    }
  }
}`,
			want: &Station{
				Name:    "US Embassy in Beijing",
				City:    "Beijing",
				State:   "Beijing",
				Country: "China",
				Location: &Location{
					Type:        "Point",
					Coordinates: []float64{116.466258, 39.954352},
				},
				Forecasts: []*Forecast{
					{
						TS:    "2019-08-05T03:00:00.000Z",
						AQIUS: 41,
						AQICN: 14,
						TP:    25,
						TPMin: 25,
						PR:    962,
						HU:    65,
						WS:    1,
						WD:    228,
						IC:    "03n",
					},
					{
						TS:    "2019-08-07T00:00:00.000Z",
						AQIUS: 68,
						AQICN: 29,
					},
				},
				Current: &Current{
					Weather: &Weather{
						TS: "2019-08-01T23:00:00.000Z",
						TP: 37,
						PR: 1007,
						HU: 14,
						WS: 1,
						WD: 110,
						IC: "01d",
					},
					Pollution: &Pollution{
						TS:     "2019-08-04T19:00:00.000Z",
						AQIUS:  70,
						MAINUS: "p2",
						AQICN:  30,
						MAINCN: "p2",
						P2: &Unit{
							CONC:  21,
							AQIUS: 70,
							AQICN: 30,
						},
						P1: &Unit{
							CONC:  30,
							AQIUS: 27,
							AQICN: 30,
						},
						O3: &Unit{
							CONC:  48,
							AQIUS: 38,
							AQICN: 30,
						},
						N2: &Unit{
							CONC:  8,
							AQIUS: 2,
							AQICN: 8,
						},
						S2: &Unit{
							CONC:  1,
							AQIUS: 1,
							AQICN: 3,
						},
						CO: &Unit{
							CONC:  0.2,
							AQIUS: 2,
							AQICN: 2,
						},
					},
				},
				History: &History{
					Weather: []*Weather{
						{
							TS: "2019-08-01T23:00:00.000Z",
							TP: 37,
							PR: 1007,
							HU: 14,
							WS: 1,
							WD: 110,
							IC: "01d",
						},
						{
							TS: "2019-08-01T04:00:00.000Z",
							TP: 31,
							PR: 1005,
							HU: 26,
							WS: 1,
							WD: 40,
							IC: "01n",
						},
					},
					Pollution: []*Pollution{
						{
							TS:     "2019-08-04T19:00:00.000Z",
							AQIUS:  70,
							MAINUS: "p2",
							AQICN:  30,
							MAINCN: "p2",
							P2: &Unit{
								CONC:  21,
								AQIUS: 70,
								AQICN: 30,
							},
							P1: &Unit{
								CONC:  30,
								AQIUS: 27,
								AQICN: 30,
							},
							O3: &Unit{
								CONC:  48,
								AQIUS: 38,
								AQICN: 30,
							},
							N2: &Unit{
								CONC:  8,
								AQIUS: 2,
								AQICN: 8,
							},
							S2: &Unit{
								CONC:  1,
								AQIUS: 1,
								AQICN: 3,
							},
							CO: &Unit{
								CONC:  0.2,
								AQIUS: 2,
								AQICN: 2,
							},
						},
						{
							TS:     "2019-08-04T18:00:00.000Z",
							AQIUS:  57,
							MAINUS: "p2",
							AQICN:  28,
							MAINCN: "o3",
							P2: &Unit{
								CONC:  15,
								AQIUS: 57,
								AQICN: 21,
							},
							P1: &Unit{
								CONC:  22,
								AQIUS: 20,
								AQICN: 22,
							},
							O3: &Unit{
								CONC:  45,
								AQIUS: 36,
								AQICN: 28,
							},
							N2: &Unit{
								CONC:  8,
								AQIUS: 2,
								AQICN: 8,
							},
							CO: &Unit{
								CONC:  0.2,
								AQIUS: 2,
								AQICN: 2,
							},
						},
					},
				},
			},
			err: nil,
		},
		{
			name: "nearest station by IP request failed",
			result: `{
  "status": "call_limit_reached",
  "data": null
}`,
			want: nil,
			err:  fmt.Errorf("unable to retrieve nearest station by IP address geolocation: %v", "call_limit_reached"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, err := client.NearestStationIP()
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

func TestNearestStationGPS(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *Station
		err    error
	}{
		{
			name: "nearest station by GPS request success",
			result: `{
  "status": "success",
  "data": {
    "name": "US Embassy in Beijing",
    "city": "Beijing",
    "state": "Beijing",
    "country": "China",
    "location": {
      "type": "Point",
      "coordinates": [
        116.466258,
        39.954352
      ]
    },
    "forecasts": [
      {
        "ts": "2019-08-05T03:00:00.000Z",
        "aqius": 41,
        "aqicn": 14,
        "tp": 25,
        "tp_min": 25,
        "pr": 962,
        "hu": 65,
        "ws": 1,
        "wd": 228,
        "ic": "03n"
      },
      {
        "ts": "2019-08-07T00:00:00.000Z",
        "aqius": 68,
        "aqicn": 29
      }
    ],
    "current": {
      "weather": {
        "ts": "2019-08-01T23:00:00.000Z",
        "tp": 37,
        "pr": 1007,
        "hu": 14,
        "ws": 1,
        "wd": 110,
        "ic": "01d"
      },
      "pollution": {
        "ts": "2019-08-04T19:00:00.000Z",
        "aqius": 70,
        "mainus": "p2",
        "aqicn": 30,
        "maincn": "p2",
        "p2": {
          "conc": 21,
          "aqius": 70,
          "aqicn": 30
        },
        "p1": {
          "conc": 30,
          "aqius": 27,
          "aqicn": 30
        },
        "o3": {
          "conc": 48,
          "aqius": 38,
          "aqicn": 30
        },
        "n2": {
          "conc": 8,
          "aqius": 2,
          "aqicn": 8
        },
        "s2": {
          "conc": 1,
          "aqius": 1,
          "aqicn": 3
        },
        "co": {
          "conc": 0.2,
          "aqius": 2,
          "aqicn": 2
        }
      }
    },
    "history": {
      "weather": [
        {
          "ts": "2019-08-01T23:00:00.000Z",
          "tp": 37,
          "pr": 1007,
          "hu": 14,
          "ws": 1,
          "wd": 110,
          "ic": "01d"
        },
        {
          "ts": "2019-08-01T04:00:00.000Z",
          "tp": 31,
          "pr": 1005,
          "hu": 26,
          "ws": 1,
          "wd": 40,
          "ic": "01n"
        }
      ],
      "pollution": [
        {
          "ts": "2019-08-04T19:00:00.000Z",
          "aqius": 70,
          "mainus": "p2",
          "aqicn": 30,
          "maincn": "p2",
          "p2": {
            "conc": 21,
            "aqius": 70,
            "aqicn": 30
          },
          "p1": {
            "conc": 30,
            "aqius": 27,
            "aqicn": 30
          },
          "o3": {
            "conc": 48,
            "aqius": 38,
            "aqicn": 30
          },
          "n2": {
            "conc": 8,
            "aqius": 2,
            "aqicn": 8
          },
          "s2": {
            "conc": 1,
            "aqius": 1,
            "aqicn": 3
          },
          "co": {
            "conc": 0.2,
            "aqius": 2,
            "aqicn": 2
          }
        },
        {
          "ts": "2019-08-04T18:00:00.000Z",
          "aqius": 57,
          "mainus": "p2",
          "aqicn": 28,
          "maincn": "o3",
          "p2": {
            "conc": 15,
            "aqius": 57,
            "aqicn": 21
          },
          "p1": {
            "conc": 22,
            "aqius": 20,
            "aqicn": 22
          },
          "o3": {
            "conc": 45,
            "aqius": 36,
            "aqicn": 28
          },
          "n2": {
            "conc": 8,
            "aqius": 2,
            "aqicn": 8
          },
          "co": {
            "conc": 0.2,
            "aqius": 2,
            "aqicn": 2
          }
        }
      ]
    }
  }
}`,
			want: &Station{
				Name:    "US Embassy in Beijing",
				City:    "Beijing",
				State:   "Beijing",
				Country: "China",
				Location: &Location{
					Type:        "Point",
					Coordinates: []float64{116.466258, 39.954352},
				},
				Forecasts: []*Forecast{
					{
						TS:    "2019-08-05T03:00:00.000Z",
						AQIUS: 41,
						AQICN: 14,
						TP:    25,
						TPMin: 25,
						PR:    962,
						HU:    65,
						WS:    1,
						WD:    228,
						IC:    "03n",
					},
					{
						TS:    "2019-08-07T00:00:00.000Z",
						AQIUS: 68,
						AQICN: 29,
					},
				},
				Current: &Current{
					Weather: &Weather{
						TS: "2019-08-01T23:00:00.000Z",
						TP: 37,
						PR: 1007,
						HU: 14,
						WS: 1,
						WD: 110,
						IC: "01d",
					},
					Pollution: &Pollution{
						TS:     "2019-08-04T19:00:00.000Z",
						AQIUS:  70,
						MAINUS: "p2",
						AQICN:  30,
						MAINCN: "p2",
						P2: &Unit{
							CONC:  21,
							AQIUS: 70,
							AQICN: 30,
						},
						P1: &Unit{
							CONC:  30,
							AQIUS: 27,
							AQICN: 30,
						},
						O3: &Unit{
							CONC:  48,
							AQIUS: 38,
							AQICN: 30,
						},
						N2: &Unit{
							CONC:  8,
							AQIUS: 2,
							AQICN: 8,
						},
						S2: &Unit{
							CONC:  1,
							AQIUS: 1,
							AQICN: 3,
						},
						CO: &Unit{
							CONC:  0.2,
							AQIUS: 2,
							AQICN: 2,
						},
					},
				},
				History: &History{
					Weather: []*Weather{
						{
							TS: "2019-08-01T23:00:00.000Z",
							TP: 37,
							PR: 1007,
							HU: 14,
							WS: 1,
							WD: 110,
							IC: "01d",
						},
						{
							TS: "2019-08-01T04:00:00.000Z",
							TP: 31,
							PR: 1005,
							HU: 26,
							WS: 1,
							WD: 40,
							IC: "01n",
						},
					},
					Pollution: []*Pollution{
						{
							TS:     "2019-08-04T19:00:00.000Z",
							AQIUS:  70,
							MAINUS: "p2",
							AQICN:  30,
							MAINCN: "p2",
							P2: &Unit{
								CONC:  21,
								AQIUS: 70,
								AQICN: 30,
							},
							P1: &Unit{
								CONC:  30,
								AQIUS: 27,
								AQICN: 30,
							},
							O3: &Unit{
								CONC:  48,
								AQIUS: 38,
								AQICN: 30,
							},
							N2: &Unit{
								CONC:  8,
								AQIUS: 2,
								AQICN: 8,
							},
							S2: &Unit{
								CONC:  1,
								AQIUS: 1,
								AQICN: 3,
							},
							CO: &Unit{
								CONC:  0.2,
								AQIUS: 2,
								AQICN: 2,
							},
						},
						{
							TS:     "2019-08-04T18:00:00.000Z",
							AQIUS:  57,
							MAINUS: "p2",
							AQICN:  28,
							MAINCN: "o3",
							P2: &Unit{
								CONC:  15,
								AQIUS: 57,
								AQICN: 21,
							},
							P1: &Unit{
								CONC:  22,
								AQIUS: 20,
								AQICN: 22,
							},
							O3: &Unit{
								CONC:  45,
								AQIUS: 36,
								AQICN: 28,
							},
							N2: &Unit{
								CONC:  8,
								AQIUS: 2,
								AQICN: 8,
							},
							CO: &Unit{
								CONC:  0.2,
								AQIUS: 2,
								AQICN: 2,
							},
						},
					},
				},
			},
			err: nil,
		},
		{
			name: "nearest station by GPS request failed",
			result: `{
  "status": "call_limit_reached",
  "data": null
}`,
			want: nil,
			err:  fmt.Errorf("unable to retrieve nearest station by GPS coordinates: %v", "call_limit_reached"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, err := client.NearestStationGPS(116.466258, 39.954352)
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
