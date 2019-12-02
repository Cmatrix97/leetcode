package problem

//Solution122 买卖股票的最佳时机II
func Solution122()  {
	prices := []int{7,1,5,3,6,4}
	println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}