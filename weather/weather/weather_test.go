package weather

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchWeatherData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)


}
