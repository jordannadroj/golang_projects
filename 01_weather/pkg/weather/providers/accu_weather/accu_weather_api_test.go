package accu_weather

import (
	"github.com/stretchr/testify/assert"
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

//
//func TestAccuWeatherAPI_Get(t *testing.T) {
//	getRequestFunc = func(url string) (*http.Response, error) {
//		return &http.Response{
//			StatusCode: http.StatusOK,
//			Body:       ioutil.NopCloser(strings.NewReader(`{"temperature": {"metric": {"value": 12.2}"}}`)),
//		}, nil
//	}
//}

// write a test for for a successful response from the API
func TestAccuWeatherAPI_Get(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"temperature": {"metric": {"value": 12.2}"}}`)),
		}, nil
	}

	response, err := fakeAccuWeatherApi.Get("Berlin")
	assert.NotNil(t, &response)
	assert.Nil(t, err)
	assert.EqualValues(t, "Berlin", &response.CityName)
	assert.EqualValues(t, 12.2, &response.Temp)
}
