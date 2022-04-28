package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	wet "weather/pkg/weather"
	"weather/pkg/weather/providers/accu_weather"
	"weather/pkg/weather/providers/open_weather_map"
	"weather/view"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	view.PromptUser()
	city := view.GetCity()

	// initialize a new NewOpenWeatherMapAPI
	openWeather := open_weather_map.NewOpenWeatherMapAPI(os.Getenv("WEATHER_APIKEY"), wet.HTTPClientReal{})
	resp, err := getWeather(openWeather, city)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)

	fmt.Println("Now with Accu Weather'\n____________________")

	// initialize a new AccuWeatherAPI
	accuWeatherAPI := accu_weather.NewAccuWeatherApi(os.Getenv("ACCUWEATHERAPIKEY"), wet.HTTPClientReal{})

	resp, err = getWeather(accuWeatherAPI, city)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

// because both API's satisfy the weather provider interface, we can pass any into this function to get weather.
func getWeather(provider wet.WeatherProvider, city string) (string, error) {
	weather, err := provider.Get(city)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("The weather in %s is %v°C\n", weather.CityName, weather.Temp), nil
}
