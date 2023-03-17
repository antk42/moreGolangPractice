package main

import "fmt"

type Engineer struct {
	Name string
}

func (e *Engineer) updateName(name string) {
	e.Name = name
}

func doStuff(e *Engineer) {
	name := "Johny Sparrow"
	defer e.updateName(name)
	fmt.Println("doing other exciting stuff")

	name = "Mr Johny Sparrow"
}

func main() {
	johny := &Engineer{
		Name: "Johny",
	}
	fmt.Printf("%+v\n", johny)
	doStuff(johny)
	fmt.Printf("%+v\n", johny)
}
