package main

import (
	"fmt"
	"log"
	"weather/view"

	"github.com/joho/godotenv"

	// can alias the package with something shorter, i.e. wet
	wet "weather/weather"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(Run())
}

func Run() string {
	view.PromptUser()
	city := view.GetCity()

	currentCity := wet.Location{CityName: city}

	response, err := currentCity.CallWeatherApi()
	if err != nil {
		log.Fatal("error calling API")
	}

	responseData, err := wet.ReadResponseBody(response)
	if err != nil {
		log.Fatal("error reading response body")
	}

	responseObject, err := wet.UnmarshalResponse(responseData)
	if err != nil {
		log.Fatal("error unmarshal response body")
	}

	return wet.Temperature(city, responseObject)

}
