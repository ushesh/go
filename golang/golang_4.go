package main

import (
	"fmt"
	"math"
)

type Vertex struct { // method will be defined on struct type
	X, Y float64
}

type Person struct {
	Name string
	Age  int
}

type MyFloat float64 // methods can be defined on any type, including built-ins

func (v *Vertex) Abs() float64 { // method defined on a struct, we use a pointer
                                 // to the struct as it's more memory efficient
                                 // and we can modify its contents
																// in case we don't send a pointer, it will work
																// anyway since we're not modifying the value
																// just using it
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) { // if we don't use a pointer here, the
                                    // function will make a copy of v and modify
																		// that instead of the value sent in from
																		// the calling function
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
type Stringer interface {
    String() string
}

This is implemented by "fmt" and used in the function below that automatically
gets used when strings are printed by "fmt" commands.
We can use the exact same command for any type and define it like below for any
formatting that is returned as a string.
We must use fmt.Sprintf and the right solution is usually to do it withing one
fmt.Sprintf statement like below

*/

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//golang.org tour - methods and interfaces
func main () {
	fmt.Println("Part 4 - methods and interfaces")
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

  // trying below just to make sure my logic was OK for the IP address exercise
	// this didn't work there as there's no fmt.Sprintf that Stringer needs
	x1 := [4]byte{8, 8, 8, 8}
	x2 := make([]byte, 0)
	for _, c := range x1 {
		x2 = append(x2, c)
	}
	fmt.Println(x2)
}
