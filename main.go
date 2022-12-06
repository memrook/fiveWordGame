package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const wordbook = "russian_nouns.txt"
const wordLen = 5

func main() {
	file, err := os.Open(wordbook)
	if err != nil {
		log.Panicln(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//var i int
	var words []string

	for scanner.Scan() {
		word := scanner.Text()
		runes := []rune(word)
		if len(runes) == wordLen {
			words = append(words, word)
		}
	}

	wordCount := len(words)
	rand.Seed(time.Now().UnixNano())
	randomNr := rand.Intn(wordCount)

	secretWord := words[randomNr]

	log.Println(wordCount, randomNr, secretWord)

	var outputWord []string
	var input string

	for fmt.Scanf("%s", &input); input != secretWord; fmt.Scanf("%s", &input) {
		input = strings.ToLower(input)
		inputRunes := []rune(input)[:wordLen]
		secretRunes := []rune(secretWord)

		log.Println("Вы ввели: ", string(inputRunes))

		for i, inputChar := range inputRunes {
			include := 0
			for _, secretChar := range secretRunes {
				if secretChar == inputChar {
					include++
				} else {
					continue
				}
			}
			switch {
			case inputChar == secretRunes[i]:
				outputWord = append(outputWord, strings.ToUpper(string(inputChar)))
			case include > 0 && inputChar != secretRunes[i]:
				outputWord = append(outputWord, string(inputChar))
			default:
				outputWord = append(outputWord, "*")
			}
		}
	}
	fmt.Println("ВЫ ПОБЕДИЛИ!!!")

}
