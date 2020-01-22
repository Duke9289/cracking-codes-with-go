package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	myMessage := "Cenoonommstmme oo snnio. s s c"
	myKey := 8

	plaintext := decryptMessage(myKey, myMessage)

	fmt.Println(plaintext + "|")
}

func decryptMessage(key int, message string) string {
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
