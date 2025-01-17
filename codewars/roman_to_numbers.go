package main

import "fmt"

func main() {
	tests := []string{"MCMXC", "MM", "MDCLXVI", "M", "CD", "XC", "XL", "I", "ABC"}

	for _, test := range tests {
		fmt.Println(convertRomanToNum(test))
	}
}

var romanNumbers = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func convertRomanToNum(roman string) int {
	prevValue := 0
	total := 0

	for i := len([]rune(roman)) - 1; i >= 0; i-- {
		char := rune(roman[i])
		currentValue, exist := romanNumbers[char]

		if !exist {
			continue
		}

		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}
	return total
}
