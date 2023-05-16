package main

import "fmt"

func main() {
	a := 3
	b := 5

	var res bool

	res = a == b
	fmt.Printf("%d es igual %d : %t \n", a, b, res)

	res = a != b
	fmt.Printf("%d es distinto %d : %t \n", a, b, res)

	res = a > b
	fmt.Printf("%d es mayor %d : %t \n", a, b, res)

	res = a < b
	fmt.Printf("%d es menor %d : %t \n", a, b, res)

	res = a >= b
	fmt.Printf("%d es mayor o igual %d : %t \n", a, b, res)

	res = a <= b
	fmt.Printf("%d es menor o igual %d : %t \n", a, b, res)
}
