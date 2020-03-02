package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	var inputFile string
	flag.StringVar(&inputFile, "i", "", "The file to read from.  Reads from STDIN if not specified")

	var outputFile string
	flag.StringVar(&outputFile, "o", "", "The file to save output to.  Prints to STDOUT if not specified")

	flag.Parse()

	input := flag.Args()
	cipher := input[0]
	var key string
	if len(input) == 2 {
		key = input[1]
	}

	var text string
	if inputFile != "" {
		fileData, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		text = string(fileData)
	} else {
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
	}

	result := ciphers[cipher](text, key, decrypt)

	if outputFile != "" {
		writeToFile(outputFile, result)
	} else {
		fmt.Println(result)
	}
}

func confirmFileWrite(outputFile string) bool {
	info, err := os.Stat(outputFile)
	if os.IsNotExist(err) || info.IsDir() {
		return true
	} else {
		fmt.Printf("File %s exists, please type 'yes' to overwrite, anything else to cancel: ", outputFile)
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		return strings.TrimRight(response, "\n") == "yes"
	}
}

func writeToFile(outputFile string, result string) {
	if confirmFileWrite(outputFile) {
		out, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating file", err)
			return
		}
		defer out.Close()
		_, err = out.WriteString(result)
		if err != nil {
			fmt.Println("Error writing to file", err)
			return
		}
	} else {
		fmt.Println("Cancelling")
	}
}
