package weather

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeatherData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`WeatherData{CityName: "berlin", Temp: 9.5}`))
	}))
	defer server.Close()

	weatherData, _ := GetWeatherData("berlin")
	if weatherData.Temp != 9.5 {
		t.Errorf("Expected 9.5 but got %v", weatherData.Temp)
	}
}
