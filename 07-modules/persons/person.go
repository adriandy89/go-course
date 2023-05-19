package persons

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Constructor(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) SetName(newName string) {
	p.name = newName
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) Print() {
	fmt.Printf("Hello %s, age: %d, ref:%p\n", p.name, p.age, &p)
}
