package main

import (
	"fmt"
	"time"
)

func helloThere(name string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("hello: %s\n", name)
}

func main() {
	go helloThere("Johny")
	fmt.Println("It should be printed first")
	time.Sleep(2 * time.Second)
}
