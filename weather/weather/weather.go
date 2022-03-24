package weather

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func PromptUser() {
	fmt.Printf("Type in a city for it's weather > ")
}

func AskAgain() {
	fmt.Printf("Another city? y/n > ")

	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	switch answer[0] {
	case 'y':
		GetCity()
	case 'Y':
		GetCity()
	case 'n':
		os.Exit(3)
	case 'N':
		os.Exit(3)
	default:
		AskAgain()
	}
}

func GetCity() {
	PromptUser()

	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = city[:len(city)-1]
	cityWeather, err := FetchWeatherData(city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cityWeather)
	AskAgain()
}

func FetchWeatherData(city string) (string, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&units=metric&&appid=%v", city, os.Getenv("WEATHER_APIKEY"))

	//response in is bytes format
	response, err := http.Get(url)
	response.Header.Add("Accept", "application/json")

	if err != nil {
		return "", err
	}

	// we then perform the conversion of our response’s body from bytes into something meaningful that can be printed out in the console.
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseObject WeatherData
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return "", err
	}
	if responseObject.Main.Temp == 0 {
		fmt.Println(responseObject.Main.Temp)
		return "", errors.New("invalid city")
	}
	return fmt.Sprintf("The current temperature in %s is %v°C", city, responseObject.Main.Temp), nil
}
