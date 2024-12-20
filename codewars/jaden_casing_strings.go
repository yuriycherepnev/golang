/*
конкатенация строк через += может быть неэффективной при работе с большими строками (каждый раз создается новая строка).
Вместо этого можно использовать strings.Builder для более эффективной работы со строками.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	jadenCase := ToJadenCase("to upper case")
	fmt.Println(jadenCase)
}

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	var builder strings.Builder
	for _, word := range words {
		builder.WriteString(strings.ToUpper(string(word[0])) + word[1:] + " ")
	}
	return strings.TrimSpace(builder.String())
}
