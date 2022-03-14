package weather

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

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
	fmt.Println(FetchWeatherData(city))
	AskAgain()
}

func FetchWeatherData(city string) string {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&units=metric&&appid=%v", city, os.Getenv("WEATHER_APIKEY"))

	//response in is bytes format
	response, err := http.Get(url)
	response.Header.Add("Accept", "application/json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// we then perform the conversion of our response’s body from bytes into something meaningful that can be printed out in the console.
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return fmt.Sprintf("The current temperature in %s is %v°C", city, responseObject.Main.Temp)
}
