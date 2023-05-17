package main

import "fmt"

func main() {
	chain := "Hello world"
	fmt.Printf("%p\n", &chain)
	fmt.Printf("Before function: %s\n", chain)
	updateValue(chain) // send a copy, not the original, not reference
	fmt.Printf("After function: %s\n", chain)
	updateValueWithRef(&chain) // send a reference to the original
	fmt.Printf("After function with pinter: %s\n", chain)
}

func updateValue(chain string) {
	fmt.Printf("%p\n", &chain)
	chain = "Hello from function"
}

// TODO - pointers *
func updateValueWithRef(chain *string) {
	fmt.Printf("%p\n", &*chain)
	*chain = "Hello from function with ref"
}

/*
	Output:
	0xc00009e230
	Before function: Hello world
	0xc00009e250
	After function: Hello world
	0xc00009e230
	After function with pinter: Hello from function with ref
*/
