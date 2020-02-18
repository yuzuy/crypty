package crypty

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const MaxRandomNumberRange = 234

// Encrypt returns the crypty hash.
func Encrypt(text string) string {
	var encryptedText string
	randNumRange := MaxRandomNumberRange
	unicodeCodePoints := []rune(text)
	rand.Seed(time.Now().UnixNano())
	for _, code := range unicodeCodePoints {
		if MaxRandomNumberRange*MaxRandomNumberRange-randNumRange*randNumRange < int(code) {
			randNumRange -= int(math.Sqrt(float64(code)))
		}
	}
	for _, code := range unicodeCodePoints {
		randNum := rand.Intn(randNumRange)
		if randNum == 0 {
			randNum = 1
		}
		encryptedText += fmt.Sprintf("%s", string(code+rune(randNum))+string(rune(randNum*randNum)))
	}
	return encryptedText
}

func Decrypt(encryptedText string) string {
	var planeText string
	words := strings.Split(encryptedText, "")
	for {
		randNum := math.Sqrt(float64([]rune(words[1])[0]))
		decryptedWord := string([]rune(words[0])[0] - rune(randNum))
		planeText += decryptedWord
		words = words[2:]
		if len(words) == 0 {
			break
		}
	}
	return planeText
}
