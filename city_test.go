package airvisual

import (
	"reflect"
	"testing"
)

func TestCities(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *Cities
	}{
		{
			name: "cities request success",
			result: `{
  "status": "success",
  "data": [
    {
      "city": "Addison"
    },
    {
      "city": "Albany"
    }
  ]
}`,
			want: &Cities{
				Status: "success",
				Data: []*CitiesData{
					{City: "Addison"},
					{City: "Albany"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.Cities("New York", "USA")
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
		})
	}
}

func TestCity(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *City
	}{
		{
			name: "city request success",
			result: `{
  "status": "success",
  "data": {
    "city": "Los Angeles",
    "state": "California",
    "country": "USA",
    "location": {
      "type": "Point",
      "coordinates": [
        -118.2417,
        34.0669
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
			want: &City{
				Status: "success",
				Data: &CityData{
					City:    "Los Angeles",
					State:   "California",
					Country: "USA",
					Location: &Location{
						Type:        "Point",
						Coordinates: []float64{-118.2417, 34.0669},
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
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.City("Los Angeles", "California", "USA")
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
		})
	}
}

func TestNearestCityIP(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *NearestCity
	}{
		{
			name: "nearest city by IP request success",
			result: `{
  "status": "success",
  "data": {
    "city": "Los Angeles",
    "state": "California",
    "country": "USA",
    "location": {
      "type": "Point",
      "coordinates": [
        -118.2417,
        34.0669
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
			want: &NearestCity{
				Status: "success",
				Data: &CityData{
					City:    "Los Angeles",
					State:   "California",
					Country: "USA",
					Location: &Location{
						Type:        "Point",
						Coordinates: []float64{-118.2417, 34.0669},
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
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.NearestCityIP()
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
		})
	}
}

func TestNearestCityGPS(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *NearestCity
	}{
		{
			name: "nearest city by IP request success",
			result: `{
  "status": "success",
  "data": {
    "city": "Los Angeles",
    "state": "California",
    "country": "USA",
    "location": {
      "type": "Point",
      "coordinates": [
        -118.2417,
        34.0669
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
			want: &NearestCity{
				Status: "success",
				Data: &CityData{
					City:    "Los Angeles",
					State:   "California",
					Country: "USA",
					Location: &Location{
						Type:        "Point",
						Coordinates: []float64{-118.2417, 34.0669},
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
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.NearestCityGPS(-118.2417, 34.0669)
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
		})
	}
}

func TestCityRanking(t *testing.T) {
	tests := []struct {
		name   string
		result string
		want   *CityRanking
	}{
		{
			name: "city ranking request success",
			result: `{
  "status": "success",
  "data": [
    {
      "city": "Portland",
      "state": "Oregon",
      "country": "USA",
      "ranking": {
        "current_aqi": 183,
        "current_aqi_cn": 154
      }
      },
      {
      "city": "Eugene",
      "state": "Oregon",
      "country": "USA",
      "ranking": {
        "current_aqi": 151,
        "current_aqi_cn": 77
      }
    }
  ]
}`,
			want: &CityRanking{
				Status: "success",
				Data: []*CityRankingData{
					{
						City:    "Portland",
						State:   "Oregon",
						Country: "USA",
						Ranking: &Ranking{
							CurrentAQI:   183,
							CurrentAQICN: 154,
						},
					},
					{
						City:    "Eugene",
						State:   "Oregon",
						Country: "USA",
						Ranking: &Ranking{
							CurrentAQI:   151,
							CurrentAQICN: 77,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, server := mockClientServer(test.result)
			defer server.Close()

			got, _ := client.CityRanking()
			want := test.want

			if !reflect.DeepEqual(want, got) {
				t.Errorf("expected %#v , got %#v", want, got)
			}
		})
	}
}
