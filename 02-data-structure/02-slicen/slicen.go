package main

import "fmt"

func main() {
	// Slicen
	slicen1 := []int{1, 2, 3}
	fmt.Println(slicen1, len(slicen1))

	// add
	slicen1 = append(slicen1, 4)
	slicen1 = append(slicen1, 5)

	fmt.Println(slicen1, len(slicen1))

	// TODO - Subslicen
	subSlicen := slicen1[:2]

	slicen1[1] = 10
	subSlicen[0] = 20

	// TODO - IMPORTANT - it's reference to the original
	fmt.Println(subSlicen) // [20 10]
	fmt.Println(slicen1)   // [20 10 3 4 5]

	// TODO - Long, Capacity, Memory Reference
	// TODO - When cap == len, duplicate cap and change memory reference
	numbers := []int{1, 2, 3}
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 4)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change
	numbers = append(numbers, 5)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change
	numbers = append(numbers, 6)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change
	numbers = append(numbers, 7)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change

}
