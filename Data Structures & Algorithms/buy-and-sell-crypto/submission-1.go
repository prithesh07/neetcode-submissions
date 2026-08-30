func maxProfit(prices []int) int {
	n := len(prices)
	left, right := 0, 0
	max_profit := 0

	for right < n {
		profit := prices[right] - prices[left]

		if profit >= 0 {
			max_profit = max(max_profit, profit)
			right++
		} else {
			left = right
		}
	}

	return max_profit
}
