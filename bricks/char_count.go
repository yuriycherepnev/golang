package main

func main() {

}

/* через слайс рун */
/* дополнительный расход памяти */
func charCount1(str string) int {
	return len([]rune(str))
}

/* через for range */
/* нет доп расходов памяти */
func charCount2(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}
