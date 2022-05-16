package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortLink(t *testing.T) {
	initialLink_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink_1 := GenerateShortLink(initialLink_1)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2 := GenerateShortLink(initialLink_2)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortLink(initialLink_3)

	assert.Equal(t, "ZV7FREWM", shortLink_1)
	assert.Equal(t, "hFYku75v", shortLink_2)
	assert.Equal(t, "AFZXB5Hs", shortLink_3)
}
