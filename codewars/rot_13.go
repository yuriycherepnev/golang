package main

import (
	"fmt"
	"unicode"
)

func main() {
	convert := Rot13("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	fmt.Println(convert)
}

func Rot13(msg string) string {
	runes := []rune(msg)

	for i, r := range runes {
		if unicode.IsLetter(r) {
			start := 'A'
			end := 'Z'
			if r >= 'a' && r <= 'z' {
				start = 'a'
				end = 'z'
			}
			convertRune := r + 13
			if convertRune > end {
				convertRune = convertRune - end + (start - 1)
			}
			runes[i] = convertRune
		} else {
			runes[i] = r
		}
	}

	return string(runes)
}
