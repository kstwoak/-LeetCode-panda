package main

import "fmt"

func main() {

	a := []int{1, 2, 5}
	b := 13

	fmt.Println(coinChange(a, b))
}

func coinChange(coins []int, amount int) int {

	// 初始化一个切片 作为dp table 默认值为 0
	dp := make([]int, amount+1)

	// 除了第一位之外都置为amount+1
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1

	}
	// 双层循环操作dp
	// 外层选择最小值
	for i := 1; i <= amount; i++ {
		// 内层选择所有值
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
