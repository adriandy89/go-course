package main

import "fmt"

func main() {
	assignation := 2
	fmt.Println(assignation)

	// assignation = assignation + 1;
	assignation++
	assignation++
	assignation++
	assignation++
	fmt.Println(assignation)
	assignation--
	assignation--
	fmt.Println(assignation)
}
