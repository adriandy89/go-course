package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	// f := n * factorial(n-1)
	// return f
	return n * factorial(n-1)
}
