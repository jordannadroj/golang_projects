package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Weather interface {
	CallWeatherApi() (*http.Response, error)
}

type Location struct {
	CityName string
}

func (city *Location) CallWeatherApi() (*http.Response, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&units=metric&&appid=%v", city.CityName, os.Getenv("WEATHER_APIKEY"))

	//response in is bytes format
	response, err := http.Get(url)
	response.Header.Add("Accept", "application/json")

	if err != nil {
		return nil, err
	}
	return response, nil
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
	var responseObject WeatherData
	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return WeatherData{}, err
	}
	return responseObject, nil
}

func Temperature(city string, responseObject WeatherData) string {
	return fmt.Sprintf("The current temperature in %s is %v°C", city, responseObject.Main.Temp)
}
