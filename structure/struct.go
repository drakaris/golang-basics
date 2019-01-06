package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func resolveVertex(v Vertex) Vertex {
	v.X = 4
	v.Y = 5

	return v
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(resolveVertex(v))
}
