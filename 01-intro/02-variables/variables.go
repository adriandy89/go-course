package main

import "fmt"

func main() {

	var name1 string
	name1 = "Name 1"

	var name2 string = "Name 2"

	name3 := "Name 3"

	fmt.Println(name1, name2, name3)

	var a int
	var b float32
	var c string
	var d bool
	fmt.Println(a, b, c, d) //  0 0  false

	const pi = 3.14
	fmt.Println(pi)
}
