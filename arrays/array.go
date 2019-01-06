package main

import "fmt"

func main() {
	var finalString [2]string

	finalString[0] = "Hello"
	finalString[1] = "World"

	fmt.Println(finalString[0], finalString[1])
	fmt.Println(finalString)

	var primes = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(primes)
}
