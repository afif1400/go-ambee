package go_ambee

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetAirQualityByCity returns the air quality of a city
func (c *AmbeeClient) GetAirQualityByCity(ctx context.Context, city string) (*apiResponse, error) {
	if city == "" {
		ErrCityEmpty := errors.New("city is empty")
		return nil, ErrCityEmpty
	}

	req, err := http.NewRequest("GET", c.BaseURL+"/by-city?city="+city, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := apiResponse{}
	_ = c.do(req, &res)

	if res.Stations == nil {
		ErrNoStations := errors.New("no stations found")
		return nil, ErrNoStations
	}

	return &res, nil
}

// GetAirQualityByPostalCode returns the air quality of a postal code
func (c *AmbeeClient) GetAirQualityByPostalCode(ctx context.Context, postalCode string, countryCode string) (*apiResponse, error) {
	if postalCode == "" {
		ErrPostalCodeEmpty := errors.New("postalCode is empty")
		return nil, ErrPostalCodeEmpty
	}

	if countryCode == "" {
		ErrCountryEmpty := errors.New("country is empty")
		return nil, ErrCountryEmpty
	}

	req, err := http.NewRequest("GET", c.BaseURL+"/by-postal-code?postalCode="+postalCode+"&countryCode="+countryCode, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := apiResponse{}
	_ = c.do(req, &res)

	if res.Stations == nil {
		ErrNoStations := errors.New("no stations found")
		return nil, ErrNoStations
	}

	return &res, nil
}

func (c *AmbeeClient) do(req *http.Request, res *apiResponse) error {
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return errors.New("status code error: " + string(rune(resp.StatusCode)))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	return nil
}
