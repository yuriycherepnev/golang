package main

func main() {

}

func containsDuplicate(nums []int) bool {
	keyNumbers := make(map[int]bool)

	for _, value := range nums {
		if keyNumbers[value] {
			return true
		}
		keyNumbers[value] = true
	}
	return false
}
