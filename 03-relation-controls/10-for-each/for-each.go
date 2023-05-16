package main

import "fmt"

func main() {
	names := [...]string{"Adrian", "Pepe", "Jose"}
	/*
		for i := 0; i < len(names); i++ {
			fmt.Println(names[i])
		}
	*/
	for index, element := range names {
		fmt.Println(index, element)
	}
	for _, element := range names {
		fmt.Println(element)
	}
	for index, _ := range names {
		fmt.Println(index)
	}
	for index := range names {
		fmt.Println(index)
	}
}
