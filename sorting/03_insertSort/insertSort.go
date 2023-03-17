package main

import "fmt"

func InsertSort(input []int) []int {
	i := 1
	for i < len(input) {
		var j = i
		for j >= 1 && input[j] < input[j-1] {
			fmt.Println("swapping")
			input[j], input[j-1] = input[j-1], input[j]
			fmt.Println(input)

			j--
		}
		i++
	}
	return input
}

func main() {
	unsortedInput := []int{5, 3, 4, 1, 2}
	sorted := InsertSort(unsortedInput)
	fmt.Println(sorted)
}
