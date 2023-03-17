package main

import (
	"fmt"
	"math/rand"
)

func main() {
	values := make(chan int)
	go CalculateValue(values)
	go CalculateValue(values)

	value := <-values
	value2 := <-values
	fmt.Println(value)
	fmt.Println(value2)
}

func CalculateValue(values chan int) {
	value := rand.Intn(20)
	fmt.Printf("Value Calculated: %d\n", value)
	values <- value
}
