package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	result := InArray(a1, a2)
	fmt.Println(result)
}

func InArray(array1 []string, array2 []string) []string {
	result := make([]string, 0)

	for _, str1 := range array1 {
		for _, str2 := range array2 {
			if strings.Contains(str2, str1) {
				result = append(result, str1)
				break
			}
		}
	}

	sort.Strings(result)

	return result
}
