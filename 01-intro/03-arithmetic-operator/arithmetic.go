package main

import "fmt"

func main() {
	a := 3
	b := 2

	result := a + b
	fmt.Println("Suma: ", result)

	result = b - a
	fmt.Println("Resta: ", result)

	result = a * b
	fmt.Println("Multiplicacion: ", result)

	result = a / b
	fmt.Println("Division: ", result) // 1

	div := 3.0 / 2.0
	fmt.Println("Division float: ", div) // 1.5

	result = a % b
	fmt.Println("Modulo: ", result)
}
