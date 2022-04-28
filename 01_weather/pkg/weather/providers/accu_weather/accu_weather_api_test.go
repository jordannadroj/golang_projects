package accu_weather

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var getRequestFunc func(url string) (*http.Response, error)

type getClientMock struct{}

func (cm getClientMock) Get(url string) (*http.Response, error) {
	response, err := getRequestFunc(url)
	return response, err
}

var fakeAccuWeatherApi = AccuWeatherAPI{
	apiKey:     "api_key",
	httpClient: getClientMock{},
}

func TestAccuWeatherAPI_Get(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"temperature": {"metric": {"value": 12.2}"}}`)),
		}, nil
	}
}
