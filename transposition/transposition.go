package transposition

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Cipher(message string, inputKey string, decryptFlag bool) string {
	var result string
	key, err := strconv.Atoi(inputKey)
	if err != nil {
		fmt.Println("No valid key supplied.  Attempting to crack.")
		return "TRANSPOSITION CRACK NOT YET IMPLEMENTED"
	}
	if decryptFlag {
		result = decrypt(key, message)
	} else {
		result = encrypt(key, message)
	}
	return result
}

func encrypt(key int, message string) string {
	ciphertext := make([]string, len(message))

	for i := 0; i < key; i++ {
		currentIndex := i

		for currentIndex < len(message) {
			ciphertext[i] += string(message[currentIndex])
			currentIndex += key
		}
	}

	return strings.Join(ciphertext, "")
}

func decrypt(key int, message string) string {
	numOfColumns := int(math.Ceil(float64(len(message)) / float64(key)))
	numOfRows := key
	numOfShadedColumns := (numOfColumns * numOfRows) - len(message)

	plaintext := make([]string, numOfColumns)
	column, row := 0, 0

	for _, runeValue := range message {
		plaintext[column] += string(runeValue)
		column += 1
		if column == numOfColumns ||
			(column == numOfColumns-1 && row >= numOfRows-numOfShadedColumns) {
			column = 0
			row += 1
		}
	}

	return strings.Join(plaintext, "")
}
