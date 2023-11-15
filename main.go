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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNr := r.Intn(wordCount)

	secretWord := words[randomNr]

	fmt.Printf(colorLime("Правила игры просты:\nЯ загадываю слово из %d букв - вы отгадываете.\n"+
		"В ответе:\n"+
		"\t-буквы в верхнем регистре - угадали в правильном месте\n"+
		"\t-буквы в нижнем регистре - они точно есть в слове, но не в этом месте\n"), wordLen)

	var outputWord []string
	var excludedChars []string
	var input string
	var count = 1

	//writer := uilive.New()       // writer for the first line
	//writer2 := writer.Newline()  // writer for the second line
	//// start listening for updates and render
	//writer.Start()

	for fmt.Scanf("%s\n", &input); input != secretWord; fmt.Scanf("%s\n", &input) {
		input = strings.ToLower(input)
		inputRunes := []rune(input)[:wordLen]
		secretRunes := []rune(secretWord)
		outputWord = nil

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
				excludedChars = appendIfNotExists(excludedChars, colorBgYellow(strings.ToUpper(string(inputChar))))
			}
		}
		fmt.Printf("\r\033[k #%d\t%v\t ❌  Исключенные символы:%v\n", count, outputWord, excludedChars)

		count++
	}

	outputWord = []string{}
	for _, v := range secretWord {
		outputWord = append(outputWord, colorLime(strings.ToUpper(string(v))))
	}

	fmt.Printf("\r\033[k #%d\t%v\t ❌  Исключенные символы:%v\n", count, outputWord, excludedChars)
	fmt.Println(colorLime("ВЫ ПОБЕДИЛИ!!!\nКол-во попыток: "), count)

}

func appendIfNotExists(origin []string, new string) []string {
	for _, char := range origin {
		if char == new {
			return origin
		}
	}
	return append(origin, new)
}
