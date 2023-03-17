package main

import "fmt"

func main() {
	HelloThere("Johny Sparrow", 42, 186)

	total, negativeTotal := AddTotal(2, 3)
	fmt.Println(total)
	fmt.Println(negativeTotal)
}

func HelloThere(name string, age int, height int) {
	fmt.Println("Hello", name)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)
}

func AddTotal(a, b int) (int, int) {
	return a + b, a - b
}
