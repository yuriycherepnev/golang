package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	result := compareSlices(a1, a2)
	fmt.Println(result)
}

func compareSlices(sliceOne, sliceTwo []string) []string {
	result := make([]string, 0)

	for _, str1 := range sliceOne {
		for _, str2 := range sliceTwo {
			if strings.Contains(str2, str1) {
				result = append(result, str1)
				break
			}
		}
	}

	sort.Strings(result)

	return result
}
