package test

type WeatherData struct {
	Temp         int
	ChanceOfRain uint
}

type FetchAPIInterface interface {
	GetWeatherData(cityName string) (WeatherData, error)
}

func main() {

}
