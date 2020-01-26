package main

import (
	"fmt"

	"github.com/Duke9289/cracking-codes-with-go/transposition"
)

func main() {
	myMessage := "Common sense is not so common."
	myKey := 8

	ciphertext := transposition.Encrypt(myKey, myMessage)

	fmt.Println(ciphertext + "|")

	myCipher := "Cenoonommstmme oo snnio. s s c"
	plaintext := transposition.Decrypt(myKey, myCipher)

	fmt.Println(plaintext + "|")
}
