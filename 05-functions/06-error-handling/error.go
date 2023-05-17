package main

import (
	"errors"
	"fmt"
)

func main() {
	res, error := div(2, 0)
	if error == nil {
		fmt.Println("Result: ", res)
	} else {
		fmt.Println("Error: ", error)
	}

}

func div(div1, div2 int) (int, error) {
	if div2 == 0 {
		return 0, errors.New("Not possible divide by '0'")
	}
	return div1 / div2, nil
}
