package main

import "fmt"

func main() {
	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "nepture"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "mercury")
	fmt.Println(planetSliceVerbose)
}
