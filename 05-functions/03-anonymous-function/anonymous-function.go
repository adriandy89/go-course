package main

import "fmt"

func main() {
	// anonymous function
	func() {
		fmt.Println("Hola desde la funcion anonima")
	}()

	myFunc := func(name string) {
		fmt.Printf("%s, Hola desde la funcion anonima en la variable\n", name)
	}

	// myFunc - alone its a memory reference

	// myFunc() - execute
	myFunc("Adrian")
}
