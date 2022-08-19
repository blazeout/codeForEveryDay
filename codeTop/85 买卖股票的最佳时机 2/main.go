package main

func maxProfit(prices []int) int {
	// 把上行的全部买下即可
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

func main() {
	// 牛逼
}
