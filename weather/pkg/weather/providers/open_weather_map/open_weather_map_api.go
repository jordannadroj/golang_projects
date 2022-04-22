package open_weather_map

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"weather/pkg/weather"
)

const BASE_URL = "https://api.openweathermap.org/data/2.5"

// the enrtre API response as a struct. I have commented out what we do not need and only included the data we care about.
type OpenWeatherMapAPIResponse struct {
	//Coord struct {
	//	Lon float64 `json:"lon"`
	//	Lat float64 `json:"lat"`
	//} `json:"coord"`
	//Weather []struct {
	//	ID          int    `json:"id"`
	//	Main        string `json:"main"`
	//	Description string `json:"description"`
	//	Icon        string `json:"icon"`
	//} `json:"weather"`
	//Base string `json:"base"`
	Main struct {
		Temp float64 `json:"temp"`
		//FeelsLike float64 `json:"feels_like"`
		//TempMin   float64 `json:"temp_min"`
		//TempMax   float64 `json:"temp_max"`
		//Pressure  int     `json:"pressure"`
		//Humidity  int     `json:"humidity"`
	} `json:"main"`
	//Visibility int `json:"visibility"`
	//Wind       struct {
	//	Speed float64 `json:"speed"`
	//	Deg   int     `json:"deg"`
	//} `json:"wind"`
	//Clouds struct {
	//	All int `json:"all"`
	//} `json:"clouds"`
	//Dt  int `json:"dt"`
	//Sys struct {
	//	Type    int    `json:"type"`
	//	ID      int    `json:"id"`
	//	Country string `json:"country"`
	//	Sunrise int    `json:"sunrise"`
	//	Sunset  int    `json:"sunset"`
	//} `json:"sys"`
	//Timezone int    `json:"timezone"`
	//ID       int    `json:"id"`
	Name string `json:"name"`
	//Cod      int    `json:"cod"`
}

type OpenWeatherMapAPI struct {
	apiKey     string
	httpClient weather.HTTPClient
}

// builds a new OpenWeatherMapAPI
func NewOpenWeatherMapAPI(apiKey string, httpClient weather.HTTPClient) *OpenWeatherMapAPI {

	return &OpenWeatherMapAPI{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (i *OpenWeatherMapAPI) Get(cityName string) (*weather.WeatherData, error) {
	url := fmt.Sprintf("%s/weather?q=%v&units=metric&&appid=%v", BASE_URL, cityName, i.apiKey)
	response, err := i.httpClient.Get(url)
	if err != nil {
		return nil, weather.ErrInternalServer
	}
	switch response.StatusCode {
	case 401:
		return nil, weather.ErrUnauthorized
	case 404:
		return nil, weather.ErrCityNotFound
	}

	parsedData, err := i.parseResponse(&response.Body)
	if err != nil {
		return nil, weather.ErrInternalServer
	}

	return i.createWeatherData(parsedData), nil
}

// Technically I could just return this and not go the extra step to return WeatherData, because I cherry-picked the data I wanted from the OpenWeatherMapAPIResponse.
//However, since WeatherData is our model, we want this data form to be consistent regardless of where we get the original data.
func (_ *OpenWeatherMapAPI) parseResponse(responseBody *io.ReadCloser) (*OpenWeatherMapAPIResponse, error) {
	responseData, err := ioutil.ReadAll(*responseBody)
	if err != nil {
		return nil, err
	}

	var responseObject OpenWeatherMapAPIResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}
	return &responseObject, nil
}

func (_ *OpenWeatherMapAPI) createWeatherData(result *OpenWeatherMapAPIResponse) *weather.WeatherData {
	return &weather.WeatherData{
		CityName: result.Name,
		Temp:     result.Main.Temp,
	}
}
