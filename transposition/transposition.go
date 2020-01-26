package transposition

import (
	"math"
	"strings"
)

func Encrypt(key int, message string) string {
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

func Decrypt(key int, message string) string {
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
