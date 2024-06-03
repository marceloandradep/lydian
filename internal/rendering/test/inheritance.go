package test

import "fmt"

type Person interface {
	GetName() string
	GetAge() string
}

type Parent struct {
	name string
}

func (p *Parent) GetName() string {
	return p.name
}

type Child struct {
	Parent
	age string
}

func (c *Child) GetAge() string {
	return fmt.Sprintf("%s is %s years old.", c.name, c.age)
}

func test() {
	var p Person
	p = &Child{
		Parent: Parent{
			name: "Marcelo",
		},
		age: "42",
	}
	println(p.GetAge())
}
