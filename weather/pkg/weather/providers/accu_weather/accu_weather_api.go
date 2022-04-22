package accu_weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather/pkg/weather"
)

// make first call to get location key, grab location key

// make another call to get weather info

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

//TODO: Refactor this into smaller functions similar to openweather map
func (i *AccuWeatherAPI) Get(cityName string) (*weather.WeatherData, error) {
	locationUrl := fmt.Sprintf("%s/locations/v1/cities/search?apikey=%v&q=%s", ACCU_URL, i.apiKey, cityName)
	locationResponse, err := http.Get(locationUrl)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	locationResponseData, err := ioutil.ReadAll(locationResponse.Body)
	var locationKeys AccuWeatherAPILocationKeyResponse
	err = json.Unmarshal(locationResponseData, &locationKeys)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	//	now use location key to call accuweather data, with first index of locationKey
	weatherUrl := fmt.Sprintf("%s/currentconditions/v1/%s?apikey=%s", ACCU_URL, locationKeys[0].Key, i.apiKey)
	weatherResponse, err := http.Get(weatherUrl)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	weatherResponseData, err := ioutil.ReadAll(weatherResponse.Body)
	var accuWeatherData AccuWeatherAPIWeatherResponse
	err = json.Unmarshal(weatherResponseData, &accuWeatherData)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	return &weather.WeatherData{
		CityName: cityName,
		Temp:     accuWeatherData[0].Temperature.Metric.Value,
	}, nil

}
