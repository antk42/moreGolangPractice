package main

import "fmt"

func main() {

	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "uranus", "nepture"}
	fmt.Println(planets)

	var planetsArray [8]string
	planetsArray[0] = "mercury"
	fmt.Println(planetsArray)
}
