package main

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	wet "weather/pkg/weather"
	"weather/view"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	view.PromptUser()
	city := view.GetCity()

	// initialize a new NewOpenWeatherMapAPI
	openWeather := wet.NewOpenWeatherMapAPI(os.Getenv("WEATHER_APIKEY"), wet.HTTPClientReal{})
	getWeather(openWeather, city)

	fmt.Println("Now with Accu Weather'\n____________________")

	// initialize a new AccuWeatherAPI
	accuWeatherAPI := wet.NewAccuWeatherApi(os.Getenv("ACCUWEATHERAPIKEY"), wet.HTTPClientReal{})

	getWeather(accuWeatherAPI, city)

}

// because both API's satisfy the wethaer provider interface, we can pass any into this function to get weather.
func getWeather(provider wet.WeatherProvider, city string) {
	weather, err := provider.Get(city)
	if err != nil {
		errors.New("invalid city, exiting")
	}
	fmt.Printf("The weather in %s is %v°C\n", weather.CityName, weather.Temp)
}
