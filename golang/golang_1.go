package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// golang.org tour - Packages, variables and functions
func main () {
	var i int
	a, b := "hello", "world"
	fmt.Println("Part 1 - Packages, variables and functions")
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)
	i = 2
	fmt.Println(i + add(45, 64))
}
