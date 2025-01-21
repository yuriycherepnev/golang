package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(numbers [10]uint) string {
	var sb strings.Builder
	for index, number := range numbers {
		if index == 0 {
			sb.WriteString("(")
		}
		if index == 3 {
			sb.WriteString(") ")
		}
		if index == 6 {
			sb.WriteString("-")
		}
		sb.WriteString(strconv.Itoa(int(number)))
	}
	return sb.String()
}
