package main

import (
	"fmt"
	"packages/animals"
	"packages/messages"
	"packages/persons"
)

func main() {
	messages.Hello()
	messages.Print()

	animalList := []animals.Animal{
		animals.Dog{Name: "Fido"},
		animals.Cat{Name: "Whiskers"},
		animals.Dog{Name: "Max"},
	}
	for _, animal := range animalList {
		//fmt.Println(animal.Speak())
		animal.Speak()
	}
	cat := animals.Cat{Name: "Kitty"}
	animals.AnimalSpeak(&cat)

	p1 := persons.Person{}
	p1.Constructor("Adrian", 30)
	fmt.Println(p1.GetName())
	p1.SetName("Name 1")
	fmt.Println(p1.GetName())
	p1.Print()
}
