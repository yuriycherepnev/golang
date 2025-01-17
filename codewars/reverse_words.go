package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(stringProcessing("reverse word aloha"))
	fmt.Println(stringProcessing("Hey fellow warriors")) // "Hey wollef sroirraw"
	fmt.Println(stringProcessing("This is a test"))      // "This is a test"
	fmt.Println(stringProcessing("This is another test"))
	fmt.Println(stringProcessing("Привет гайз , как делишки сегодня?"))
}

func stringProcessing(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len([]rune(word)) >= 5 {
			words[i] = reverseWord(word)
		}
	}
	return strings.Join(words, " ")
}

func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len([]rune(word))-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
