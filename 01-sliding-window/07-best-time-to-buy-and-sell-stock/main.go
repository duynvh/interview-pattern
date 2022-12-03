package main

func maxProfit(stockPrices []int) int {
	left, right, maxProfit := 0, 1, 0
	for right < len(stockPrices) {
		if stockPrices[left] < stockPrices[right] {
			profit := stockPrices[right] - stockPrices[left]
			maxProfit = max(maxProfit, profit)
		} else {
			left = right
		}
		right++
	}

	return maxProfit
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
