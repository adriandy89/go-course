package main

import "fmt"

func main() {
	numbers1 := make([]int, 0)
	fmt.Println(numbers1, len(numbers1), cap(numbers1))
	numbers2 := make([]int, 3, 3)
	fmt.Println(numbers2, len(numbers2), cap(numbers2))
	numbers3 := make([]int, 3)
	fmt.Println(numbers3, len(numbers3), cap(numbers3))
	numbers4 := make([]int, 0, 3)
	fmt.Println(numbers4, len(numbers4), cap(numbers4))

	numbers2[0] = 5
	numbers2[1] = 5
	numbers2[2] = 5
	numbers2 = append(numbers2, 4)
	fmt.Println(numbers2, len(numbers2), cap(numbers2))

}
