package airvisual

// Location is an object containing location information
type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Forecast is an object containing forecast information
type Forecast struct {
	TS    string  `json:"ts"`               // timestamp
	AQIUS int     `json:"aqius"`            // AQI value based on US EPA standard
	AQICN int     `json:"aqicn"`            // AQI value based on China MEP standard
	TP    float64 `json:"tp,omitempty"`     // temperature in Celsius
	TPMin float64 `json:"tp_min,omitempty"` // minimum temperature in Celsius
	PR    float64 `json:"pr,omitempty"`     // atmospheric pressure in hPa
	HU    float64 `json:"hu,omitempty"`     // humidity %
	WS    float64 `json:"ws,omitempty"`     // wind speed (m/s)
	WD    float64 `json:"wd,omitempty"`     // wind direction, as an angle of 360° (N=0, E=90, S=180, W=270)
	IC    string  `json:"ic,omitempty"`     // weather icon code, see below for icon index
}

// Weather contains weather information
type Weather struct {
	TS string  `json:"ts"`
	TP float64 `json:"tp"`
	PR float64 `json:"pr"`
	HU float64 `json:"hu"`
	WS float64 `json:"ws"`
	WD float64 `json:"wd"`
	IC string  `json:"ic"`
}

// Pollution contains pollution information
type Pollution struct {
	TS     string `json:"ts"`
	AQIUS  int    `json:"aqius"`
	MAINUS string `json:"mainus"` // main pollutant for US AQI
	AQICN  int    `json:"aqicn"`
	MAINCN string `json:"maincn"` // main pollutant for Chinese AQI
	// pollutant details, concentration and appropriate AQIs
	P2 *Unit `json:"p2,omitempty"`
	P1 *Unit `json:"p1,omitempty"`
	O3 *Unit `json:"o3,omitempty"`
	N2 *Unit `json:"n2,omitempty"`
	S2 *Unit `json:"s2,omitempty"`
	CO *Unit `json:"co,omitempty"`
}

// Current is an object containing weather and pollution current information
type Current struct {
	Weather   *Weather   `json:"weather"`
	Pollution *Pollution `json:"pollution"`
}

// History is an object containing weather and pollution history information
type History struct {
	Weather   []*Weather   `json:"weather"`
	Pollution []*Pollution `json:"pollution"`
}

// Ranking is an object containing current AQI for specific country
type Ranking struct {
	CurrentAQI   int `json:"current_aqi"`
	CurrentAQICN int `json:"current_aqi_cn"`
}

//Unit is a polution unit
type Unit struct {
	CONC  float64 `json:"conc"`
	AQIUS int     `json:"aqius"`
	AQICN int     `json:"aqicn"`
}
