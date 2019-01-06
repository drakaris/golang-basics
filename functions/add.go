package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("The sum is: ", add(10, 20))
}
