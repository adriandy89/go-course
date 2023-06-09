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
	fmt.Println("TEST")
	// TODO - IMPORTANT - it's reference to the original
	fmt.Println(slicen1)   // [20 10 3 4 5]
	fmt.Println(subSlicen) // [20 10]

	fmt.Println("NUMBERS")
	// TODO - Long, Capacity, Memory Reference
	// TODO - When cap == len, duplicate cap and change memory reference
	numbers := []int{1, 2, 3}
	subNumbers := numbers[:2]
	fmt.Println(numbers)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers)
	fmt.Println(subNumbers)
	fmt.Printf("Sub - Len: %d, Capac: %v, Memory Ref: %p \n", len(subNumbers), cap(subNumbers), subNumbers)
	numbers[1] = 5
	fmt.Println(numbers)
	fmt.Println(subNumbers)
	numbers = append(numbers, 4)
	numbers[1] = 10
	fmt.Println(numbers)
	fmt.Println(subNumbers)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change
	fmt.Println(subNumbers)
	// TODO - different, don't change, memory ref change on last append.
	fmt.Printf("Sub - Len: %d, Capac: %v, Memory Ref: %p \n", len(subNumbers), cap(subNumbers), subNumbers)
	numbers = append(numbers, 5)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change
	numbers = append(numbers, 6)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change
	numbers = append(numbers, 7)
	fmt.Printf("Len: %d, Capac: %v, Memory Ref: %p \n", len(numbers), cap(numbers), numbers) // change

}
