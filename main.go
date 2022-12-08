package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const wordbook = "wordbook.txt"
const wordLen = 5

var colorLime = color.New(color.FgHiGreen).SprintFunc()
var colorYellow = color.New(color.FgYellow).SprintFunc()
var colorBgYellow = color.New(color.BgYellow).SprintFunc()

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

	//log.Println(wordCount, randomNr, secretWord)
	fmt.Printf(colorLime("Правила игры просты:\nЯ загадываю слово из %d букв - вы отгадываете.\n"+
		"В ответе:\n"+
		"\t-буквы в верхнем регистре - угадали в правильном месте\n"+
		"\t-буквы в нижнем регистре - они точно есть в слове, но не в этом месте\n"), wordLen)

	var outputWord []string
	var excludedChars []string
	var input string

	for fmt.Scanf("%s", &input); input != secretWord; fmt.Scanf("%s", &input) {
		input = strings.ToLower(input)
		inputRunes := []rune(input)[:wordLen]
		secretRunes := []rune(secretWord)
		outputWord = nil

		//fmt.Println("Вы ввели: ", string(inputRunes))

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
				outputWord = append(outputWord, colorLime(strings.ToUpper(string(inputChar))))
			case include > 0 && inputChar != secretRunes[i]:
				outputWord = append(outputWord, colorYellow(string(inputChar)))
			default:
				outputWord = append(outputWord, "*")
				excludedChars = appendIfNotExists(excludedChars, colorBgYellow(string(inputChar)))
				//excludedChars = append(excludedChars, colorBgYellow(string(inputChar)))
			}
		}
		fmt.Printf("%v\t\t ❌  Исключенные символы:%v\n", outputWord, excludedChars)
	}
	fmt.Println("ВЫ ПОБЕДИЛИ!!!")

}

func appendIfNotExists(origin []string, new string) []string {
	for _, char := range origin {
		if char == new {
			return origin
		}
	}
	return append(origin, new)
}
