package main

import "fmt"

func main() {
	test()        // []
	test(4, 5, 3) // [4 5 3]
	message, total := sum("Suma", 5, 3, 9, 7, 5, 12)
	fmt.Printf("La %s es igual a %d\n", message, total)
}

func test(numbers ...int) {
	fmt.Println(numbers)
}

func sum(operation string, numbers ...int) (string, int) {
	var total int
	for _, v := range numbers {
		total += v
	}
	return operation, total
}
