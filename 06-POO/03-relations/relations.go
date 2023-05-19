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

// TODO - relation
type Employed struct {
	person Person
	salary float32
}

type Group struct {
	employed []Employed
	group    string
}

func main() {

	// TODO - inheritance
	employed1 := Employed{salary: 50}
	employed1.person.name = "H1"
	employed1.person.age = 20
	employed1.person.print()
	fmt.Println(employed1)
	employed1.salary = 1500
	fmt.Println(employed1)

	p1 := Person{"Adrian", 30}
	p1.print()
	e2 := Employed{
		person: p1,
		salary: 23,
	}
	fmt.Println(e2)
	p1.age = 31 // dont
	fmt.Println(p1)
	e2.salary += 50
	fmt.Println(e2)

	group1 := Group{
		group:    "Group 1",
		employed: []Employed{employed1, e2},
	}
	fmt.Println(group1)
}

/*
	Hello H1, age: 20, ref:0xc000012028
	{{H1 20} 50}
	{{H1 20} 1500}
	Hello Adrian, age: 30, ref:0xc000012038
	{{Adrian 30} 23}
	{Adrian 31}
	{{Adrian 30} 73}
*/
