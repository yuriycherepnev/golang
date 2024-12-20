package main

import (
	"fmt"
	"strings"
)

func WordCount(input string) map[string]int {
	words := strings.Split(input, " ")

	//string - ключ / int - значение
	mapWords := make(map[string]int)

	for _, word := range words {
		// Убираем пунктуацию и приводим к нижнему регистру
		word = strings.Trim(word, ",.!?")
		word = strings.ToLower(word)

		// Увеличиваем счетчик
		mapWords[word]++
	}

	return mapWords
}

func main() {
	fmt.Println(WordCount("Go is fun, and Go is powerful!"))
}
