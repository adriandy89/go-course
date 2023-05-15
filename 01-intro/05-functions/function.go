package main

import "fmt"

func main() {
	hello()
	sum := sum(3, 5)
	fmt.Printf("Suma = %d \n", sum)
}

func hello() {
	fmt.Println("Hello")
}

func sum(a, b int) int {
	return a + b
}
