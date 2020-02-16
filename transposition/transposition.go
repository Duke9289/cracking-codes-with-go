package transposition

import (
	"math"
	"strconv"
	"strings"
)

func Cipher(message string, inputKey string, mode string) string {
	var result string
	key, err := strconv.Atoi(inputKey)
	if err != nil {
		return "The key for a transposition cipher must be an integer"
	}
	switch mode {
	case "encrypt":
		result = encrypt(key, message)
	case "decrypt":
		result = decrypt(key, message)
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
