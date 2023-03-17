package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Engineer) Print() {
	fmt.Println("Engineer Information")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
}

func (e *Engineer) UpdateAge() {
	e.Age += 1
}

func (e *Engineer) GetProjectPriority() string {
	return e.Project.Priority
}

func main() {
	engineer := Engineer{
		Name: "Elliot",
		Age:  27,
		Project: Project{
			Name:         "become senior golang developer",
			Priority:     "High",
			Technologies: []string{"Golang"},
		},
	}
	engineer.UpdateAge()
	engineer.Print()

	fmt.Println(engineer.GetProjectPriority())
}
