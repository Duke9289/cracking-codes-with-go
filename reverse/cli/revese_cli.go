package main

import (
	"fmt"

	"github.com/Duke9289/cracking-codes-with-go/reverse"
)

func main() {
	message := "Three can keep a secret, if two of them are dead."
	fmt.Println(reverse.Reverse(message))
}
