package main

import (
	"fmt"

	"github.com/Duke9289/cracking-codes-with-go/caesar"
)

func main() {
	message := "This is my secret message"
	key := 13

	mode := "encrypt"

	translated := caesar.Encrypt(key, message, mode)
	fmt.Println("Encrypted message: ", translated)

	caesar.Crack(translated)
}
