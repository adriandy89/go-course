package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var expression string

	fmt.Print("==> ")
	fmt.Scanln(&expression)
	res := sum(expression)

	fmt.Println(res)
}

func sum(expression string) int {
	data := strings.Split(expression, "+")
	var result int

	for _, v := range data {
		num, error := strconv.Atoi(v)
		if error != nil {
			fmt.Println(error)
		} else {
			result += num
		}
	}
	return result
}
