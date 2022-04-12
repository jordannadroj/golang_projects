package weather

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetWeatherData(t *testing.T) {
	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *WeatherData
		expectedErr      error
	}{
		{
			"returns weather data for berlin",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"coord":{"lon":13.4105,"lat":52.5244},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":12.84,"feels_like":11.28,"temp_min":11.65,"temp_max":14.45,"pressure":1008,"humidity":42},"visibility":10000,"wind":{"speed":4.12,"deg":270},"clouds":{"all":0},"dt":1649690828,"sys":{"type":2,"id":2011538,"country":"DE","sunrise":1649650668,"sunset":1649699799},"timezone":7200,"id":2950159,"name":"Berlin","cod":200}`))
			})),
			&WeatherData{
				CityName: "Berlin",
				Temp:     12.84,
			},
			nil,
		},
		{
			"returns weather data for new york",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"coord":{"lon":-74.006,"lat":40.7143},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"base":"stations","main":{"temp":14.84,"feels_like":13.37,"temp_min":11.08,"temp_max":18.12,"pressure":1018,"humidity":38},"visibility":10000,"wind":{"speed":3.58,"deg":179,"gust":5.36},"clouds":{"all":0},"dt":1649706788,"sys":{"type":2,"id":2039034,"country":"US","sunrise":1649672595,"sunset":1649719824},"timezone":-14400,"id":5128581,"name":"New York","cod":200}`))
			})),
			&WeatherData{
				CityName: "New York",
				Temp:     14.84,
			},
			nil,
		},
		{
			"returns 404 error when using an unknown location",
			httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"cod":"404","message":"city not found"}`))
			})),
			nil,
			ErrBadRequest,
		},
		{
			name: "returns an error when sending nothing",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			})),
			expectedResponse: nil,
			expectedErr:      ErrBadRequest,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, err := GetWeatherData(tc.server.URL, "berlin")
			if !reflect.DeepEqual(resp, tc.expectedResponse) {
				t.Errorf("expectec %v, but got %v", tc.expectedResponse, resp)
			}
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected %v, but got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestTemperature(t *testing.T) {
	weatherData := WeatherData{
		CityName: "Berlin",
		Temp:     5.43,
	}
	expect := fmt.Sprintf("The current temperature in %s is %v°C", weatherData.CityName, weatherData.Temp)

	actual := Temperature(&weatherData)

	if expect != actual {
		t.Errorf("Expected %q, but got %q", expect, actual)
	}
}
