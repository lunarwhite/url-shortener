package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
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

// Hashing initialUrl + userId url with sha256
func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()          // Derive a big integer number from the hash bytes generated during the hasing
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber))) // apply base58  on the derived big integer value
	return finalString[:8]                                                   // Pick the first 8 characters
}
