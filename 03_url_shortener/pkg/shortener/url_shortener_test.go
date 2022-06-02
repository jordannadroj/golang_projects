package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortLink(t *testing.T) {
	initialLink := "https://github.com/jordannadroj/52_in_52/tree/main/03_url_shortener"
	shortLink, err := GenerateShortLink(initialLink)

	assert.Equal(t, "EkpxsNpT", shortLink)
	assert.NoError(t, err)
}
