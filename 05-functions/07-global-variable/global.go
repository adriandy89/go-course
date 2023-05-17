package main

import "fmt"

var global string

func main() {
	global = "Hello from Main"
	fmt.Println(global)
	defer func1() // execute at the end, after main()
	func2()
}

func func1() {
	global = "Hello from func1"
	fmt.Println(global)
}

func func2() {
	global = "Hello from func2"
	fmt.Println(global)
}
