package weather

import "net/http"

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

type HTTPClientReal struct{}

func (_ *HTTPClientReal) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}
