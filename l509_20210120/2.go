package main

import "fmt"

func main() {
	// for i := 0; i < 20; i++ {
	fmt.Println(fid(20))

	// }
}

/*
时间复杂度O（n）
*/
func fid(n int) int {
	if 0 == n {
		return 0
	}
	memo := make([]int, n+1)
	return helper(memo, n)
}

func helper(memo []int, n int) int {
	if 1 == n || 2 == n {
		return 1
	}

	//运算过的就直接返回
	if memo[n] != 0 {
		return memo[n]
	}
	fmt.Println(n)
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}
