package caesar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Duke9289/cracking-codes-with-go/consts"
)

func Cipher(message string, inputKey string, mode string) string {
	var result string
	key, err := strconv.Atoi(inputKey)
	if err != nil {
		return "The key for a caesar cipher must be an integer"
	}
	switch mode {
	case "encrypt", "decrypt":
		result = encrypt(key, message, mode)
	case "crack":
		result = crack(message)
	}

	return result
}

func encrypt(key int, message string, mode string) string {

	translated := ""

	for _, runeValue := range message {
		if strings.ContainsRune(consts.SYMBOLS, runeValue) {
			var translatedIndex int
			symbolIndex := strings.Index(consts.SYMBOLS, string(runeValue))
			if mode == "encrypt" {
				translatedIndex = symbolIndex + key
			} else {
				translatedIndex = symbolIndex - key
			}

			if translatedIndex >= len(consts.SYMBOLS) {
				translatedIndex = translatedIndex - len(consts.SYMBOLS)
			} else if translatedIndex < 0 {
				translatedIndex = translatedIndex + len(consts.SYMBOLS)
			}
			translated += string(consts.SYMBOLS[translatedIndex])
		} else {
			translated += string(runeValue)
		}
	}
	return translated
}

func crack(message string) string {
	var result string

	for i := 0; i < len(consts.SYMBOLS); i++ {
		translated := ""
		for _, runeValue := range message {
			if strings.ContainsRune(consts.SYMBOLS, runeValue) {
				var translatedIndex int
				symbolIndex := strings.Index(consts.SYMBOLS, string(runeValue))
				translatedIndex = symbolIndex - i

				if translatedIndex < 0 {
					translatedIndex = translatedIndex + len(consts.SYMBOLS)
				}
				translated += string(consts.SYMBOLS[translatedIndex])
			} else {
				translated += string(runeValue)
			}
		}
		result += fmt.Sprintf("Key %d: %s\n", i, translated)
	}
	return result
}
