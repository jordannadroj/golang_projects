package weather

import "errors"

type WeatherData struct {
	CityName string
	Temp     float64
}

var ErrCityNotFound = errors.New("city not found")
var ErrBadRequest = errors.New("bad request")
var ErrInternalServer = errors.New("internal server error")
