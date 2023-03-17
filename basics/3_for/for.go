package main

import "fmt"

func main() {
	ages := []int{42, 28, 32, 56, 67}

	for id, value := range ages {
		fmt.Println(id, value)
	}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}
	fmt.Println("======")

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}

	for {
		fmt.Println("hello")
		break
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
