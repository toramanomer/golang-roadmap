package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func swap(v *Vertex) {
	v.X, v.Y = v.Y, v.X
}

func main() {
	swap(&v1)
	fmt.Println("v1.X =", v1.X, "v1.Y =", v1.Y)
}
