package main

import "fmt"

func main() {
	fmt.Println(fid(20))

}

/*
时间复杂度O（n）
*/
func fid(n int) int {
	if 0 == n {
		return 0
	}
	if 1 == n || 2 == n {
		return 1
	}
	// fmt.Println(n)
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		fmt.Println(dp[i])
	}
	return dp[n]

}
