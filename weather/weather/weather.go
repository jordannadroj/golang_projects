package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherData struct {
	CityName string
	Temp     float64
}

type FetchAPI interface {
	GetWeatherData(cityName string) (WeatherData, error)
}

func GetWeatherData(cityName string) (WeatherData, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&units=metric&&appid=%v", cityName, os.Getenv("WEATHER_APIKEY"))

	//response in is bytes format
	response, err := http.Get(url)
	response.Header.Add("Accept", "application/json")

	if err != nil {
		return WeatherData{}, err
	}

	responseData, err := ReadResponseBody(response)
	if err != nil {
		return WeatherData{}, err
	}
	weatherData, err := UnmarshalResponse(responseData)
	if err != nil {
		return WeatherData{}, err
	}
	return weatherData, nil
}

// perform the conversion of our response’s body from bytes into something meaningful that can be printed out in the console.
func ReadResponseBody(response *http.Response) ([]byte, error) {
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func UnmarshalResponse(responseData []byte) (WeatherData, error) {
	var responseObject WeatherDataResponse // the big struct
	var weatherData WeatherData            // what we want
	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return WeatherData{}, err
	}
	weatherData.CityName = responseObject.Name
	weatherData.Temp = responseObject.Main.Temp
	return weatherData, nil
}

func Temperature(weatherData WeatherData) string {
	return fmt.Sprintf("The current temperature in %s is %v°C", weatherData.CityName, weatherData.Temp)
}
