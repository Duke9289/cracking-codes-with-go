package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Duke9289/cracking-codes-with-go/caesar"
	"github.com/Duke9289/cracking-codes-with-go/reverse"
	"github.com/Duke9289/cracking-codes-with-go/transposition"
)

var ciphers = map[string]func(string, string, bool) string{
	"caesar":        caesar.Cipher,
	"reverse":       reverse.Cipher,
	"transposition": transposition.Cipher,
}

func main() {
	var decrypt bool
	flag.BoolVar(&decrypt, "d", false, "Decrypt the message rather than encrypt it")

	flag.Parse()

	input := flag.Args()
	cipher := input[0]
	var key string
	if len(input) == 2 {
		key = input[1]
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(ciphers[cipher](text, key, decrypt))
}
