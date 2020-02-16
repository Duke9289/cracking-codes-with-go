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

/*
<cipher name>
encrypt|decrypt|crack -e|d|c
<key>
-i|input <filename> ? if null, stdin
-o|output <filename> ? if null, stdout
*/

var ciphers = map[string]func(string, string, string) string{
	"caesar":        caesar.Cipher,
	"reverse":       reverse.Cipher,
	"transposition": transposition.Cipher,
}

func main() {
	var cipher string
	flag.StringVar(&cipher, "c", "", "The cipher type")

	var mode string
	flag.StringVar(&mode, "m", "encrypt", "Whether to encrypt, decrypt, or crack")

	var key string
	flag.StringVar(&key, "k", "", "The cipher key")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(ciphers[cipher](text, key, mode))
}
