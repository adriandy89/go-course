package main

import "fmt"

func main() {
	fmt.Println(!false)
	fmt.Println(false && true)
	fmt.Println(false || true)

	res := 2 == 2 || !(4 > 3)
	fmt.Println(res)
	res = 2 == 2 && !(4 > 3)

	fmt.Println(res)
}
