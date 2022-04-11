package weather

import (
	"errors"
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
			"returns weather data",
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
			name: "returns an error",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			})),
			expectedResponse: nil,
			expectedErr:      errors.New("unexpected end of JSON input"),
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
