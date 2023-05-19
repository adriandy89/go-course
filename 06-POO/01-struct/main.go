package main

import "fmt"

// TODO - Model
type Person struct {
	name string
	age  int
}

// TODO - method
func (p *Person) print() {
	fmt.Printf("Hello %s, age: %d, ref:%p\n", p.name, p.age, &p)
}

func (p *Person) editName(newName string) {
	p.name = newName
}

// TODO - inheritance
type Employed struct {
	Person
	salary float32
}

func main() {
	p1 := Person{"Adrian", 30}
	p1.print()

	p2 := Person{
		name: "Name 2",
		age:  32,
	}
	p2.print()
	p2.age += 5000
	p2.editName("Immortal")
	p2.print()

	// TODO - inheritance
	employed1 := Employed{salary: 50}
	employed1.name = "H1"
	employed1.age = 20
	fmt.Println(employed1)
	employed1.print()
	fmt.Println(employed1.name)
}

/*
	Output:
	Hello Adrian, age: 30, ref:0xc000012028
	Hello Name 2, age: 32, ref:0xc000012038
	Hello Immortal, age: 5032, ref:0xc000012040
	{{H1 20} 50}
	Hello H1, age: 20, ref:0xc000012048
*/
