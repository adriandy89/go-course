package main

import "fmt"

func main() {
	if name, age := "Adriano", 30; name == "Adrian" {
		fmt.Println("Hola ", name)
	} else {
		fmt.Printf("Hola %s, edad: %d\n", name, age)
	}

	users := make(map[string]string)

	users["Adrian"] = "adrian@gmail.com"
	users["Pepe"] = "pepe@gmail.com"

	email1, ok1 := users["Adrian"]
	email2, ok2 := users["FALSE"]
	fmt.Println(email1, ok1) // adrian@gmail.com true
	fmt.Println(email2, ok2) //  false

	if email, ok := users["Adrian"]; ok {
		fmt.Println(email)
	} else {
		fmt.Println("No exist")
	}
}
