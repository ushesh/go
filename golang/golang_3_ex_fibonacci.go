package main

import "fmt"

/* Problem statement
 * Implement a fibonacci function that returns a function (a closure) that
 * returns successive fibonacci numbers.
*/
func fibonacci() func() int {
	f, g := 0, 1
	return func() int {
		f, g = g, f + g
		return f
	}
}

func main() {
	fibo := fibonacci() // initialize values, this is always the lines before
	                    // the closure function .i.e. f, g := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(fibo()) // fibo variable remembers the values within the
		                    // closure function that's attached to it
	}
}