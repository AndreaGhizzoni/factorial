package main

import "fmt"

func F(n int) int {
	if n <= 1 {
		return 1
	}

	N := n
	factOfN := 2
	for i := 1; i <= n-2; i++ {
		factOfN = N * (n - i)
		N = factOfN
	}
	return factOfN
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d)= %d\n", i, F(i))
	}
}
