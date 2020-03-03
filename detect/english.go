package detect

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var englishWords map[string]bool = loadDictionary()

var defaultEnglishWordPercent = 20.0
var defaultEnglishLetterPercent = 85.0

func IsEnglish(message string) bool {
	wordsMatch := getEnglishCount(message)*100 >= defaultEnglishWordPercent
	numLetters := len(removeNonLetters(message))
	messageLettersPercentage := float64(numLetters) / float64(len(message)) * 100
	lettersMatch := messageLettersPercentage >= defaultEnglishLetterPercent
	return (wordsMatch && lettersMatch)
}

func IsEnglishCustom(message string, wordPercent float64, letterPercent float64) bool {
	wordsMatch := getEnglishCount(message)*100 >= wordPercent
	numLetters := len(removeNonLetters(message))
	messageLettersPercentage := float64(numLetters) / float64(len(message)) * 100
	lettersMatch := messageLettersPercentage >= letterPercent
	return (wordsMatch && lettersMatch)
}

func loadDictionary() map[string]bool {
	// currently uses words_alpha.txt from github.com/dwyl/english-words
	// note: on linux, had to parse file to unix line ends instead of dos line ends
	data, err := ioutil.ReadFile("words_alpha.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}

	m := make(map[string]bool)
	for _, word := range strings.Split(string(data), "\n") {
		m[word] = true
	}
	return m
}

func removeNonLetters(message string) string {
	reg, err := regexp.Compile("[^a-zA-Z\\s]+")
	if err != nil {
		log.Fatal("Regex error:", err)
	}
	return reg.ReplaceAllString(message, "")
}

func getEnglishCount(message string) float64 {
	message = removeNonLetters(strings.ToLower(message))
	possibleWords := strings.Split(message, " ")

	if len(possibleWords) == 0 {
		return 0.0
	}

	matches := 0
	for _, word := range possibleWords {
		//fmt.Printf("in map, %s is %t\n", word, englishWords[word])
		if englishWords[word] {
			matches++
		}
	}

	return float64(matches) / float64(len(possibleWords))
}
