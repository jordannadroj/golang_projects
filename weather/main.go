package main

import (
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

	wet.GetCity()
}
