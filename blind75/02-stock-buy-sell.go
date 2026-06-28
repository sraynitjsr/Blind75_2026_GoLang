package blind75

import "fmt"

func maxProfit(prices []int) int {
	profit, minPrice := 0, prices[0]
	for price := range prices {
		minPrice = min(minPrice, price)
		profit = max(profit, price-minPrice)
	}
	return profit
}

func TwoStockBuySell() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println("Maximum Profit is", maxProfit(prices))
}

// Done
