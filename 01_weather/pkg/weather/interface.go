package weather

type WeatherProvider interface {
	Get(cityName string) (*WeatherData, error)
}
