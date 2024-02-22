package main

import "fmt" // for testing

// iterative
func sum_to_n_a(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// recursive
func sum_to_n_b(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum_to_n_b(n-1)
}

// formula
func sum_to_n_c(n int) int {
	return (n * (n + 1)) / 2
}

// for testing via `go run main.go`
func main() {
	fmt.Println(sum_to_n_a(5)) // 15
	fmt.Println(sum_to_n_b(5)) // 15
	fmt.Println(sum_to_n_c(5)) // 15
}
