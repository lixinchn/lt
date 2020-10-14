package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			if prices[i]-min > profit {
				profit = prices[i] - min
			}
		} else {
			min = prices[i]
		}
	}
	return profit
}
