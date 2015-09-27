package main

import (
	"fmt"
	"time"
	"io"
	"strings"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 4)
	for {
		n, err := r.Read(b) // func (T) Read(b []byte) (n int, err error)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
