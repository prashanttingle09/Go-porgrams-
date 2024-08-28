package main

import "fmt"

func main() {
	fmt.Println("fib sequence:", fib(9))
}

// to find next sequence not list itreative approach
func fib(n int) int {
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}

// recursive method
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
