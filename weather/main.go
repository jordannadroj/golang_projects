package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	// can alias the package with something shorter, i.e. wet
	wet "weather/weather"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	city := wet.GetCity()

	weather := wet.FetchWeatherData(city)

	fmt.Println(weather)
}
