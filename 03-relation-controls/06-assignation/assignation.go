package main

import "fmt"

func main() {
	assignation := 2
	fmt.Println(assignation)

	// assignation = assignation + 2;
	assignation += 8
	fmt.Println(assignation)
	assignation -= 4
	fmt.Println(assignation)
	assignation *= 5
	fmt.Println(assignation)
	assignation /= 4
	fmt.Println(assignation)
	assignation %= 3
	fmt.Println(assignation)

}
