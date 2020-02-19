package main

import (
	"errors"
	"testing"
)

func TestCrypty(t *testing.T) {
	exampleText := "hello"
	if err := test(exampleText); err != nil {
		t.Fatal()
	}
	exampleText = "∆å˜˚®å"
	if err := test(exampleText); err != nil {
		t.Fatal()
	}
	exampleText = "∆ƒøˆ´®¨∑øˆ∆çßµªˆ¢ºª"
	if err := test(exampleText); err != nil {
		t.Fatal()
	}
}

func test(text string) error {
	for i := 0; i < 1000; i++ {
		decryptedText := Decrypt(Encrypt(text))
		if text != decryptedText {
			return errors.New("")
		}
	}
	return nil
}
