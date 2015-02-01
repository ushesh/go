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
	
	var e = []int{1, 2, 4, 8, 16, 32}
	
	for index, value := range(e) {
		fmt.Printf("2 raise to %d is %d\n", index, value)
	}
	
    var m map[string]Vertex
	m = make(map[string]Vertex) // index is string, value is struct Vertex
							    // no length of map specified, just index type
							    // and value type
	m["Juniper"] = Vertex{1, 2}
	fmt.Println(m)

	var n = map[string]Vertex{ // map literal
		"BellLabs": Vertex{40, 74,}, // comma is needed inside the Vertex as
		                             // it’s a composite literal
		"Google":   Vertex{37, 122,},
	}
	fmt.Println(n)

	var o = map[string]Vertex{
		"MS":  {40, 74,}, // if the top level type is just the type name, it
		                  // can be ommitted (only in literals, not assignment
		                  // during program flow .i.e. o["FB"] = {0, 0} is not
		                  // legal)
		"IBM": {37, 122,},
	}
	fmt.Println(o)

	m["Cisco"] = Vertex{7, 8} // inserting elements into a map
	fmt.Println(m)

	fmt.Println(m["Juniper"]) // directly accessing values in a map

	value, present := m["Cisco"] // test if a value is present or not
	fmt.Println(value, present)

	_, brocade_present := m["Brocade"]
	if brocade_present {
		fmt.Println("Yes, it's here")
	} else {
		fmt.Println("Never heard of it")
	}

	fmt.Println(o)
	delete(o, "MS") // delete an element from a map using it's key
	fmt.Println(o)
}
