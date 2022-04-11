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

func GetWeatherData(url, cityName string) (*WeatherData, error) {
	api := fmt.Sprintf("%s/weather?q=%v&units=metric&&appid=%v", url, cityName, os.Getenv("WEATHER_APIKEY"))

	//response in is bytes format
	response, err := http.Get(api)
	response.Header.Add("Accept", "application/json")

	if err != nil {
		return nil, err
	}

	responseData, err := readResponseBody(response)
	if err != nil {
		return nil, err
	}
	weatherData, err := unmarshalResponse(responseData)
	if err != nil {
		return nil, err
	}
	return weatherData, nil
}

// perform the conversion of our response’s body from bytes into something meaningful that can be printed out in the console.
func readResponseBody(response *http.Response) ([]byte, error) {
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

func unmarshalResponse(responseData []byte) (*WeatherData, error) {
	var responseObject WeatherDataResponse // the big struct
	var weatherData = &WeatherData{}       // what we want
	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}
	weatherData.CityName = responseObject.Name
	weatherData.Temp = responseObject.Main.Temp
	return weatherData, nil
}

func Temperature(weatherData *WeatherData) string {
	return fmt.Sprintf("The current temperature in %s is %v°C", weatherData.CityName, weatherData.Temp)
}
