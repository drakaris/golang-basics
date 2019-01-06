package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 8; i++ {
		sum += i
	}

	fmt.Println("Sum is:", sum)
}
