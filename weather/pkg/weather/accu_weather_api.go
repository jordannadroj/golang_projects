package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	httpClient HTTPClient
}

func NewAccuWeatherApi(apiKey string, httpClient HTTPClient) *AccuWeatherAPI {
	return &AccuWeatherAPI{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (i *AccuWeatherAPI) Get(cityName string) (*WeatherData, error) {
	locationUrl := fmt.Sprintf("%s/locations/v1/cities/search?apikey=%v&q=%s", ACCU_URL, i.apiKey, cityName)
	locationResponse, err := http.Get(locationUrl)
	if err != nil {
		return nil, ErrInternalServer
	}

	locationResponseData, err := ioutil.ReadAll(locationResponse.Body)
	var locationKeys AccuWeatherAPILocationKeyResponse
	err = json.Unmarshal(locationResponseData, &locationKeys)
	if err != nil {
		return nil, ErrInternalServer
	}

	//	now use location key to call accuweather data, with first index of locationKey
	weatherUrl := fmt.Sprintf("%s/currentconditions/v1/%s?apikey=%s", ACCU_URL, locationKeys[0].Key, i.apiKey)
	weatherResponse, err := http.Get(weatherUrl)
	if err != nil {
		return nil, ErrInternalServer
	}

	weatherResponseData, err := ioutil.ReadAll(weatherResponse.Body)
	var accuWeatherData AccuWeatherAPIWeatherResponse
	err = json.Unmarshal(weatherResponseData, &accuWeatherData)
	if err != nil {
		return nil, ErrInternalServer
	}

	return &WeatherData{
		CityName: cityName,
		Temp:     accuWeatherData[0].Temperature.Metric.Value,
	}, nil

}
