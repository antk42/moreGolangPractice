package main

import "fmt"

func BubbleSort(input []int) []int {
	// n is the number of items in our list
	n := len(input)

	swapped := true

	for swapped {
		swapped = false
		// iterate through all of the elements in our list
		for i := 0; i < n-1; i++ {
			// if the current element is greater than the next
			// element, swap them
			if input[i] > input[i+1] {
				fmt.Println("swapping")
				input[i], input[i+1] = input[i+1], input[i]
				fmt.Println(input)
				swapped = true
			}
		}
	}
	return input
}

func main() {
	unsortedInput := []int{5, 3, 4, 1, 2}
	sorted := BubbleSort(unsortedInput)
	fmt.Println(sorted)
}
