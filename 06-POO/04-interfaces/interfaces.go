package main

import "fmt"

type Animal interface {
	Speak() string
}
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow! My name is " + c.Name
}

func animalSpeak(animal Animal) {
	fmt.Println(animal.Speak())
}

func main() {
	animals := []Animal{
		Dog{Name: "Fido"},
		Cat{Name: "Whiskers"},
		Dog{Name: "Max"},
	}
	for _, animal := range animals {
		//fmt.Println(animal.Speak())
		animalSpeak(animal)
	}
	cat := Cat{Name: "Kitty"}
	animalSpeak(&cat)
}
