package main

import (
	"fmt"
)

// Doak's workshop - day 2
func main() {
	// 1st exercise
	m := map[string]string{
		"First": "BE",
		"Second": "CCNA",
		"Third": "CCNP",
		"Fourth": "MS",
	}
	if _, present := m["Fifth"]; present {
		fmt.Println("Went beyond grad school")
	} else {
		fmt.Println("Enough money spent")
	}
}