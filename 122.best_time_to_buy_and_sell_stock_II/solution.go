package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	profit := 0
	tempProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > min {
			if prices[i]-min >= tempProfit {
				tempProfit = prices[i] - min
			} else {
				profit += tempProfit
				tempProfit = 0
				min = prices[i]
			}
		} else {
			min = prices[i]
			profit += tempProfit
			tempProfit = 0
		}
	}
	profit += tempProfit
	return profit
}
