package main

import "fmt"

func main() {
	Johny := &Engineer{
		Name: "Johny",
		Age:  42,
	}
	fmt.Println(Johny)

	Johny.UpdateAge()
	fmt.Println(Johny)

	Johny.UpdateName()
	fmt.Println(Johny)

	UpdateAge(Johny)
	fmt.Println(Johny)
}

type Engineer struct {
	Name string
	Age  int
}

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func UpdateAge(e *Engineer) {
	e.Age += 1
}

func (e *Engineer) UpdateName() {
	e.Name = "new name"
	fmt.Println(e)
}
