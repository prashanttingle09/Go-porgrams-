package main

import "fmt"

func main() {
	fmt.Println("factorail is:", fact(5))
	fmt.Println("factorail using recurvie is:", fact2(5))
}

func fact(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f = f * i
	}
	return f
}

func fact2(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact2(n-1)
}
