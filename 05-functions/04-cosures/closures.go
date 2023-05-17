package main

import (
	"fmt"
	"strings"
)

func main() {
	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola Mundo "))
	// anonymous function
	// func() {
	// 	fmt.Println("Hola desde la funcion anonima")
	// }()

	// myFunc := func(name string) {
	// 	fmt.Printf("%s, Hola desde la funcion anonima en la variable\n", name)
	// }

	// // myFunc - alone its a memory reference

	// // myFunc() - execute
	// myFunc("Adrian")
}

func repeat(n int) func(cadena string) string {

	// this func remember the father value -> n
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}

}
