package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	// golang mentioned to convert the e to float64 but doesn't say why
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
	  z := x / 2
	  for i := 0; i < 10; i++ {
		  z = z - ((z * z - x)/(2 * z)) // square root using Newton's
		  							                // formula
	  }
		return z, nil
  } else {
	  return 0, ErrNegativeSqrt(x)
  }
}

func main() {
	if v, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
	if v, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
