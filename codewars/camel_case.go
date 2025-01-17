package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(toCamelCase("the-stealth-warrior"))     // theStealthWarrior
	fmt.Println(toCamelCase("The_Stealth_Warrior"))     // TheStealthWarrior
	fmt.Println(toCamelCase("The_Stealth-Warrior"))     // TheStealthWarrior
	fmt.Println(toCamelCase("snake_case-example_here")) // snakeCaseExampleHere
}

func toCamelCase(str string) string {
	result := make([]rune, 0, len(str))
	capitalize := false

	for i, value := range str {
		if value == '-' || value == '_' {
			capitalize = true
			continue
		}
		if i != 0 && capitalize {
			result = append(result, unicode.ToUpper(value))
		} else {
			result = append(result, value)
		}
		capitalize = false
	}
	return string(result)
}
