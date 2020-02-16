package caesar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Duke9289/cracking-codes-with-go/consts"
)

func Cipher(message string, inputKey string, mode bool) string {
	key, err := strconv.Atoi(inputKey)
	if err != nil {
		fmt.Println("No valid key entered, attempting to crack.")
		return crack(message)
	}
	return encrypt(key, message, mode)
}

func encrypt(key int, message string, decrypt bool) string {

	translated := ""

	for _, runeValue := range message {
		if strings.ContainsRune(consts.SYMBOLS, runeValue) {
			var translatedIndex int
			symbolIndex := strings.Index(consts.SYMBOLS, string(runeValue))
			if decrypt {
				translatedIndex = symbolIndex - key
			} else {
				translatedIndex = symbolIndex + key
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
		result += fmt.Sprintf("Key %d: %s", i, translated)
	}
	return result
}
