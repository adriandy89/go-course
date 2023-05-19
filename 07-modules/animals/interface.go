package animals

import "fmt"

type Animal interface {
	Speak() string
}

func AnimalSpeak(animal Animal) {
	fmt.Println(animal.Speak())
}
