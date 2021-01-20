package main

import "fmt"

func main() {
	fmt.Println(fid(20))

}

func fid(n int) int {
	if 0 == n {
		return 0
	}
	if 1 == n || 2 == n {
		return 1
	}
	var prev, curr = 1, 1
	for i := 3; i <= n; i++ {
		var sum = prev + curr
		prev = curr
		curr = sum
	}
	return curr
}
