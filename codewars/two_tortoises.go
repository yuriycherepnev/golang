package main

import (
	"fmt"
)

func main() {
	result := Race(720, 850, 70)
	fmt.Println(result)
}

func Race(v1, v2, g int) [3]int {
	if v2 < v1 {
		return [3]int{-1, -1, -1}
	}
	totalSeconds := int((float64(g) / (float64(v2) - float64(v1))) * 3600)

	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return [3]int{hours, minutes, seconds}
}
