package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

/*
Here userId is added to prevent providing similar shortened urls to separate users in case they want to shorten exact same link, it's a design decision, so some implementations do this differently.
*/

func GenerateShortLink(initialLink string) string {
	randomId := uuid.New()
	urlHashBytes := sha256Of(initialLink + randomId.String())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
