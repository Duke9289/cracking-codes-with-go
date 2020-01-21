package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "guv6Jv6Jz!J6rp5r7Jzr66ntr"
	SYMBOLS := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890 !?."

	for i := 0; i < len(SYMBOLS); i++ {
		translated := ""
		for _, runeValue := range message {
			if strings.ContainsRune(SYMBOLS, runeValue) {
				var translatedIndex int
				symbolIndex := strings.Index(SYMBOLS, string(runeValue))
				translatedIndex = symbolIndex - i

				if translatedIndex < 0 {
					translatedIndex = translatedIndex + len(SYMBOLS)
				}
				translated += string(SYMBOLS[translatedIndex])
			} else {
				translated += string(runeValue)
			}
		}
		fmt.Printf("Key %d: %s\n", i, translated)

	}
}
