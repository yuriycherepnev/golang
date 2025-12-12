package find_k_closest_elements

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	closest := findClosestElements2(array, 3, 999)
	fmt.Println(closest)
	fmt.Println(array)
}

func findClosestElements2(arr []int, k, x int) []int {
	for len(arr) > k {
		if mod(arr[0]-x) > mod(arr[len(arr)-1]-x) {
			arr = arr[1:]
		} else {
			arr = arr[:len(arr)-1]
		}
	}
	return arr
}

func mod(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
