package open_weather_map

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"weather/pkg/weather"
)

// this is not thread safe
// race condition
var getRequestFunc func(url string) (*http.Response, error)

// satisfies the httpInterface. So I can mock Get
type getClientMock struct {
	StatusCode int
	Body       io.Reader
}

func (cm getClientMock) Get(url string) (*http.Response, error) {
	response, err := getRequestFunc(url)
	return response, err
}

var fakeOpenWeatherApi = OpenWeatherMapAPI{
	apiKey: "any",
	httpClient: getClientMock{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(`{"main":{"temp":12.25},"name":"Berlin"}`)),
	},
}

func TestOpenWeatherMapAPI_Get(t *testing.T) {
	// this is the mock of the response I want
	// use unique instances of getClientMock
	//getRequestFunc = func(url string) (*http.Response, error) {
	//	return &http.Response{
	//		StatusCode: http.StatusOK,
	//		Body:       ioutil.NopCloser(strings.NewReader(`{"main":{"temp":12.25},"name":"Berlin"}`)),
	//	}, nil
	//}

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

func TestOpenWeatherMapAPI_GetInvalidAPIKey(t *testing.T) {
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
