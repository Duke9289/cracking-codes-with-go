package caesar

import (
	"fmt"
	"strings"

	"github.com/Duke9289/cracking-codes-with-go/consts"
)

func Encrypt(key int, message string, mode string) string {

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

func Crack(message string) {

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
		fmt.Printf("Key %d: %s\n", i, translated)
	}
}
