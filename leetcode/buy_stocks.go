// [7,1,5,3,6,4]
package main

import "fmt"

func main() {
	stocks := []int{1, 1, 1, 1, 1, 10, 2, 2, 2, 2, 8}

	fmt.Println(maxProfit(stocks))
}

func maxProfit(prices []int) int {
	profit := 0
	buy := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			buy = prices[i]
			profit += prices[i+1] - buy
		}
	}
	return profit
}
