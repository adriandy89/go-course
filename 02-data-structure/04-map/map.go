package main

import "fmt"

func main() {
	days := make(map[int]string)
	fmt.Println(days)

	days[1] = "Dia 1"
	days[2] = "Dia 2"
	days[3] = "Dia 3"

	fmt.Println(days)

	days[2] = "vacio"
	fmt.Println(days[2])
	fmt.Println(days)

	delete(days, 2)
	fmt.Println(days)

	// TODO - allow len, but not cap
	fmt.Println(len(days))

	students := make(map[string][]int)
	students["group1"] = []int{14, 15, 16}
	students["group2"] = []int{17, 18, 19}
	fmt.Println(students)
	fmt.Println(students["group1"][1])

}
