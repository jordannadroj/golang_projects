package weather

import "net/http"

type HTTPClientReal struct{}

//main http interface
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

func (_ HTTPClientReal) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

//can use another tool such as curl.
