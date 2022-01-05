package go_ambee

type Station struct {
	CO          float64 `json:"CO"`
	NO2         float64 `json:"NO2"`
	OZONE       float64 `json:"OZONE"`
	PM10        float64 `json:"PM10"`
	PM25        float64 `json:"PM25"`
	CountryCode string  `json:"countryCode"`
	Division    string  `json:"division"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	PostalCode  string  `json:"postalCode"`
	City        string  `json:"city"`
	Place       string  `json:"placeName"`
	State       string  `json:"state"`
	UpdatedAt   string  `json:"updatedAt"`
	AQI         float64 `json:"AQI"`
	AqiInfo     AqiInfo
}

type AqiInfo struct {
	Pollutant     string  `json:"pollutant"`
	Concentration float64 `json:"concentration"`
	Category      string  `json:"category"`
}
