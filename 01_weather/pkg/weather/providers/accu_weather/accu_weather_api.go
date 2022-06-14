package accu_weather

import (
	"encoding/json"
	"fmt"
	"github.com/jordannadroj/52_in_52/01_weather/pkg/weather"
	"io"
	"io/ioutil"
	"net/http"
)

const ACCU_URL = "https://dataservice.accuweather.com"

type AccuWeatherAPILocationKeyResponse []struct {
	Key string `json:"Key"`
}

type AccuWeatherAPIWeatherResponse []struct {
	Temperature struct {
		Metric struct {
			Value float64 `json:"Value"`
		} `json:"Metric"`
	} `json:"Temperature"`
}

type AccuWeatherAPI struct {
	apiKey     string
	httpClient weather.HTTPClient
}

func NewAccuWeatherApi(apiKey string, httpClient weather.HTTPClient) *AccuWeatherAPI {
	return &AccuWeatherAPI{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (i *AccuWeatherAPI) Get(cityName string) (*weather.WeatherData, error) {
	// With AccuWeather must first get location key of given city
	locationKey, err := i.getLocationKey(cityName)
	if err != nil {
		return nil, err
	}

	//	now use location key to call accuweather data, with first index of locationKey
	weatherUrl := fmt.Sprintf("%s/currentconditions/v1/%s?apikey=%s", ACCU_URL, locationKey, i.apiKey)

	weatherResponse, err := http.Get(weatherUrl)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	if weatherResponse.StatusCode == 401 {
		return nil, weather.ErrUnauthorized
	}

	accuWeatherData, err := i.parseWeatherResponse(&weatherResponse.Body)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	return i.createWeatherData(accuWeatherData, cityName), nil
}

func (i *AccuWeatherAPI) getLocationKey(cityName string) (string, error) {
	locationUrl := fmt.Sprintf("%s/locations/v1/cities/search?apikey=%v&q=%s", ACCU_URL, i.apiKey, cityName)
	locationResponse, err := i.httpClient.Get(locationUrl)
	if err != nil {
		return "", weather.ErrInternalServer
	}
	if locationResponse.StatusCode == 401 {
		return "", weather.ErrUnauthorized
	}

	locationKeys, err := i.parseLocationResponse(&locationResponse.Body)

	if err != nil {
		return "", weather.ErrInternalServer
	}
	return locationKeys[0].Key, nil

}

func (_ *AccuWeatherAPI) parseLocationResponse(responseBody *io.ReadCloser) (AccuWeatherAPILocationKeyResponse, error) {
	locationResponseData, err := ioutil.ReadAll(*responseBody)
	var locationKeys AccuWeatherAPILocationKeyResponse
	err = json.Unmarshal(locationResponseData, &locationKeys)
	if err != nil {
		return nil, weather.ErrInternalServer
	}
	return locationKeys, nil
}

func (_ *AccuWeatherAPI) parseWeatherResponse(responseBody *io.ReadCloser) (AccuWeatherAPIWeatherResponse, error) {
	weatherResponseData, err := ioutil.ReadAll(*responseBody)
	var accuWeatherData AccuWeatherAPIWeatherResponse
	err = json.Unmarshal(weatherResponseData, &accuWeatherData)
	if err != nil {
		return nil, weather.ErrInternalServer
	}
	return accuWeatherData, nil
}

func (_ *AccuWeatherAPI) createWeatherData(accuWeatherData AccuWeatherAPIWeatherResponse, cityName string) *weather.WeatherData {
	return &weather.WeatherData{
		CityName: cityName,
		Temp:     accuWeatherData[0].Temperature.Metric.Value,
	}
}
