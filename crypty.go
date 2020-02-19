package crypty

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// maxRandomNumberRange(=234) is limit of the random number to use crypty
// If a random number is greater than 234, happen bugs with a limit of unicode
const maxRandomNumberRange = 234

// Encrypt generates crypty crypto
func Encrypt(text string) string {
	var encryptedText string
	randNumRange := maxRandomNumberRange
	runes := []rune(text)
	rand.Seed(time.Now().UnixNano())
	for _, code := range runes {
		if maxRandomNumberRange*maxRandomNumberRange-randNumRange*randNumRange < int(code) {
			randNumRange -= int(math.Sqrt(float64(code)))
		}
	}
	for _, code := range runes {
		randNum := rand.Intn(randNumRange)
		if randNum == 0 {
			randNum = 1
		}
		encryptedText += fmt.Sprintf("%s", string(code+rune(randNum))+string(rune(randNum*randNum)))
	}
	return encryptedText
}

// Decrypt decrypts text was encrypted with crypty
func Decrypt(encryptedText string) (string, error) {
	var plainText string
	words := strings.Split(encryptedText, "")
	for {
		randNum := math.Sqrt(float64([]rune(words[1])[0]))
		decryptedWord := string([]rune(words[0])[0] - rune(randNum))
		plainText += decryptedWord
		words = words[2:]
		if len(words) == 0 {
			break
		}
	}
	return plainText, nil
}
