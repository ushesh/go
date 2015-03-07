package main

import (
	"fmt"
	"runtime"
)

func short_summer(n int) int {
	var sum int = 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func long_summer(n int) int {
	var sum int = 1
	for sum < n { // this is the equivalent of a while loop
		sum += sum
	}
	return sum
}

func square_rt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		z = z - ((z * z - x)/(2 * z)) // square root using Newton's
									  // formula
	}
	return z
}

//golang.org tour - Flow control statements (for, if, else, switch)
func main() {
	fmt.Println("Part 2 - Flow control statements")

	defer_var := 7.0
	defer fmt.Println(square_rt(defer_var))

	fmt.Println(short_summer(25))
	fmt.Println(long_summer(200))

	if short_summer(25) < long_summer(200) {
		fmt.Println("Barely made it")
	} else {
		fmt.Println("Big enough")
	}

	fmt.Println(square_rt(1024))

	os := runtime.GOOS
	switch os {
		case "darwin":
			fmt.Println("Running on OS X")
		case "linux":
			fmt.Println("Running on Linux")
		default:
			fmt.Printf("Unknown platform - %s", os)
	}

	defer_var = 9.0

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
