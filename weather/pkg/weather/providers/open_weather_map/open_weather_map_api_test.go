package open_weather_map

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"weather/pkg/weather"
)

var getRequestFunc func(url string) (*http.Response, error)

type getClientMock struct{}

func (cm getClientMock) Get(url string) (*http.Response, error) {
	response, err := getRequestFunc(url)
	return response, err
}

var fakeOpenWeatherApi = OpenWeatherMapAPI{
	apiKey:     "any",
	httpClient: getClientMock{},
}

func TestOpenWeatherMapAPI_Get(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"main":{"temp":12.25},"name":"Berlin"}`)),
		}, nil
	}

	response, err := fakeOpenWeatherApi.Get("Berlin")
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, "Berlin", response.CityName)
	assert.EqualValues(t, 12.25, response.Temp)

}

func TestOpenWeatherMapAPI_GetWithInvalidCity(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(strings.NewReader(`{"cod":"404","message":"city not found"}`)),
		}, nil
	}

	response, err := fakeOpenWeatherApi.Get("Invalid_city")
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, weather.ErrCityNotFound, err)
}

func TestOpenWeatherMapAPI_GetInavlidAPIKey(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"cod":"401","message":"unauthorized"}`)),
		}, nil
	}

	response, err := fakeOpenWeatherApi.Get("Berlin")
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.EqualValues(t, weather.ErrUnauthorized, err)
}
