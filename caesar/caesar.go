package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "This is my secret message"
	key := 13

	mode := "encrypt"

	SYMBOLS := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890 !?."

	translated := ""

	for _, runeValue := range message {
		if strings.ContainsRune(SYMBOLS, runeValue) {
			var translatedIndex int
			symbolIndex := strings.Index(SYMBOLS, string(runeValue))
			if mode == "encrypt" {
				translatedIndex = symbolIndex + key
			} else {
				translatedIndex = symbolIndex - key
			}

			if translatedIndex >= len(SYMBOLS) {
				translatedIndex = translatedIndex - len(SYMBOLS)
			} else if translatedIndex < 0 {
				translatedIndex = translatedIndex + len(SYMBOLS)
			}
			translated += string(SYMBOLS[translatedIndex])
		} else {
			translated += string(runeValue)
		}
	}
	fmt.Println(translated)
}
