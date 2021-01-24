package main

func main() {

}

func coinChange(coins []int, amount int) int {

	dp := func(n int) int {
		if 0 == n {
			return 0
		}
		if n < 0 {
			return -1
		}
		for _, coin := range coins {
			subproblem := dp(n - coin)
		}
		return
	}()

	return dp(amount)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
