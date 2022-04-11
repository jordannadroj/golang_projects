package main

import (
	"fmt"
	"log"
	"weather/view"

	"github.com/joho/godotenv"

	// can alias the package with something shorter, i.e. wet
	wet "weather/weather"
)

const URL = "https://api.openweathermap.org/data/2.5"

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

	weatherData, err := wet.GetWeatherData(URL, city)
	if err != nil {
		log.Fatal("error retrieving weather data")
	}

	return wet.Temperature(weatherData)

}
