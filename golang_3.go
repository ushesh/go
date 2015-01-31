package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

//golang.org tour - pointers, structs, slices, maps
func main() {
	fmt.Println("Part 3 - More data structures")
	var v Vertex
	v.X = 1
	v.Y = 5
	fmt.Println(v)
	
	w := Vertex{3, 4}
	fmt.Println(w)
	
	u := &v // pointer indirection is transparent for structs
	fmt.Println(u.X)
	
	var a [2]string // array
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	
	b := make([]int, 5) // zero'd slice with len(b)=5
	fmt.Println(b)
	
	c := make([]int, 0, 10) // len(b)=0, cap(b)=10
	
	var d []int // nil slice
	
	fmt.Println(len(b), len(c), len(d))
	fmt.Println(cap(b), cap(c), cap(d))
	
	d = append(d, 2)
	fmt.Println(d)
}
