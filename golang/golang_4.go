package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { // method defined on a struct, we use a pointer
                                 // to the struct as it's more memory efficient
                                 // and we can modify its contents
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

//golang.org tour - methods and interfaces
func main () {
	fmt.Println("Part 4 - methods and interfaces")
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}