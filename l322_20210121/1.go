package main

func main() {

}

func coinChange(coins []int, amount int) int {

	if 0 == amount {
		return 0
	}
	if amount < 0 {
		return -1
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
