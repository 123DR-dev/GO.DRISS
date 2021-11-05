package main

import (
	"fmt"
)

func fact(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func factr(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factr(n-1)
	}
}

func main() {
	fmt.Println(fact(5))
	fmt.Println(factr(4))
}
