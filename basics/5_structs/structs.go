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
	fmt.Printf("%+v\n", engineer)
	fmt.Println(engineer.Project.Name)
}
