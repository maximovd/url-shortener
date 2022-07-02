package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// Encrypt input string into sha256 hash
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// Encode hash into base58 string
func base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	econded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	return string(econded)
}

// Generate short link from initial Link
func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
