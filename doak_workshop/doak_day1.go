package main

import (
	"fmt"
)

// Doak's workshop - day 1
func main() {
	// 1st exercise
	var x int = 7
	color := "yellow"
	fmt.Println(x, color)

	x = 10
	fmt.Println(x)

	// 2nd exercise
	fruits := make([]string, 0)
	fruits = append(fruits, "apple", "orange", "mango")
	fmt.Println(fruits[1])

	californiaFruits := fruits[0:2]
	fmt.Println(californiaFruits)

	californiaFruits[0] = "pear" // changing the slice also changes the original
	                             // that the slice was made from
	fmt.Println(californiaFruits, fruits)

	// 3rd exercise
	deviceToType := make(map[string]string)
	deviceToType["bbo1.atl12"] = "juniper"
	deviceToType["bx01.fra03"] = "ciscoxr"
	fmt.Printf("%s\n", deviceToType["bx01.fra03"])

	// 4th excerise
	deviceToTypeLiteral := map[string]string{ // literals don't need make
		"bbo1.atl12": "juniper",
	    "bx01.fra03": "ciscoxr",
	}

	fruitsLiteral := []string{"apple", "orange", "mango"}

	fmt.Println(deviceToTypeLiteral, fruitsLiteral)

	// 5th exercise
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
    }

    // 6th exercise
    names := []string{"bob", "ann", "marie", "steve"}
    for _, name := range names {
		fmt.Println(name)
	}

	nameMap := map[int]string{
		0: "john",
		1: "mary",
		2: "stevie",
	}
	for key, value := range nameMap {
		fmt.Printf("%d: %s\n", key, value)
	}
}