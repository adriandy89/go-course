package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		if i == 8 {
			fmt.Println("break")
			break
		}
		if i == 5 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(i)
	}
	/*
	   0
	   1
	   2
	   3
	   4
	   continue
	   6
	   7
	   break
	*/
}
