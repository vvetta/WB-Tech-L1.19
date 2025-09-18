package main

import "fmt"

func main() {
	var inputWord string
	fmt.Print("Введите слово, которое хотите перевернуть: ")
	fmt.Scan(&inputWord)

	flipedWord := string(FlipWord([]rune(inputWord)))
	fmt.Printf("Перевернутое слово: %s", flipedWord)
}

func FlipWord(word []rune) []rune {
	var flipedWord []rune

	for i := len(word) - 1; i >= 0; i-- {
		flipedWord = append(flipedWord, word[i])
	}

	return flipedWord
}
