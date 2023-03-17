package main

import "fmt"

func RecursiveBubbleSort(input []int, size int) []int {
	if size == 1 {
		return input
	}

	var i = 0
	for i < size-1 {
		if input[i] > input[i+1] {
			fmt.Println("swapping")
			input[i], input[i+1] = input[i+1], input[i]
			fmt.Println(input)
		}
		i++
	}
	RecursiveBubbleSort(input, size-1)
	return input
}

func main() {
	unsortedInput := []int{5, 3, 4, 1, 2}
	sorted := RecursiveBubbleSort(unsortedInput, len(unsortedInput))
	fmt.Println(RecursiveBubbleSort(sorted, len(sorted)))
}
