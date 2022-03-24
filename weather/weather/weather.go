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

func Run() {
	PromptUser()
	city := GetCity()

	response, err := CallWeatherApi(city)
	if err != nil {
		log.Fatal("error calling API")
	}

	responseData, err := ReadResponseBody(response)
	if err != nil {
		log.Fatal("error reading response body")
	}

	responseObject, err := UnmarshalResponse(responseData)
	if err != nil {
		log.Fatal("error unmarshal response body")
	}

	fmt.Println(Temperature(city, responseObject))
	PromptUserAgain()

}

func PromptUser() {
	fmt.Printf("Type in a city for it's weather > ")
}

func PromptUserAgain() {
	fmt.Printf("Type in a city for it's weather or press enter to exit > ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	if len(response) == 1 {
		os.Exit(2)
	}

}

func GetCity() string {
	reader := bufio.NewReader(os.Stdin)
	city, _ := reader.ReadString('\n')
	city = city[:len(city)-1]
	return city
}

func CallWeatherApi(city string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&units=metric&&appid=%v", city, os.Getenv("WEATHER_APIKEY"))

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
