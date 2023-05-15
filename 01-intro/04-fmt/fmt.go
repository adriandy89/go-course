package main

import "fmt"

func main() {
	name := "Nombre"
	age := 30

	fmt.Printf("Hello. I am %s, %d years old \n", name, age)
	fmt.Printf("Hello. I am %v, %v years old \n", name, age)

	message := fmt.Sprintf("Hello. I am %s, %d years old", name, age)
	fmt.Println(message)

	fmt.Printf("name: %T \n", name) // string

	fmt.Print("Ingrese nombre: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello. I am %s, %d years old \n", name, age)

}
