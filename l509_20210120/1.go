package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 20; i++ {
		fmt.Println(fid(i))

	}
}

/*
时间复杂度O（2n方）
*/
func fid(n int) int {
	if 0 == n {
		return 0
	}

	if 1 == n || 2 == n {
		return 1
	}

	return fid(n-1) + fid(n-2)
}
