package main

import (
	"fmt"
	"unicode"
)

func main() {
	wave := wave(" ab  c")
	fmt.Println(wave)
}

func wave(words string) []string {
	upperWords := make([]string, 0, len([]rune(words)))
	for index, char := range words {

		if unicode.IsLetter(char) {
			runes := []rune(words)
			runes[index] = unicode.ToUpper(runes[index])
			upperWords = append(upperWords, string(runes))
		}
	}
	return upperWords
}
