package weather

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFetchWeatherData(t *testing.T) {
	//server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)
	testCases := []struct {
		name    string
		city    string
		jordan  func()
		wantErr bool
	}{
		{
			"passes",
			"berlin",
			func() {
				fmt.Println("HI")
			},
			false,
		},
		{
			"error",
			"berlins",
			func() {},
			true,
		},
	}
	for _, tc := range testCases {

		tc.jordan()

		_, err := FetchWeatherData(tc.city)
		if tc.wantErr {
			require.Error(t, err, tc.name)
		} else {
			require.Nil(t, err, tc.name)
		}
	}
}
