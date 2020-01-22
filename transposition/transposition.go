package main

import (
	"fmt"
	"strings"
)

func main() {
	myMessage := "Common sense is not so common."
	myKey := 8

	ciphertext := encryptMessage(myKey, myMessage)

	fmt.Println(ciphertext + "|")
}

func encryptMessage(key int, message string) string {
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
