package messages

import "fmt"

func Hello() {
	fmt.Println("hello from package messages")
}

func privateFunction() {
	fmt.Println("hello from private function")
}

func Print() {
	privateFunction()
}
