package main

import "fmt"

func main() {
	numbers := 102350
	quantity := 0
	for numbers > 0 {
		numbers /= 10
		quantity++
	}
	fmt.Println("Cantidad de digitos: ", quantity)

	for i := 0; i <= 100; i++ {
		if i%20 == 0 {
			fmt.Println(i)

		}
	}
}
