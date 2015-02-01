package main

import (
	"strings"
	"fmt"
)

/* Problem statement
 * Implement WordCount. It should return a map of the counts of each “word” in
 * the string s.
*/
func WordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s) // returns an array of every white space
	                           // separated word in the string
	for _, word := range words { // check every word in the array
		_, present := count[word] // see if word is already present in map
		if present {
			count[word]++ // if present, increment the counter
		} else {
			count[word] = 1 // if not, add it to map with count=1
		}
	}
	return count
}

func main() {
	fmt.Println(WordCount("I am learning Go"))
	fmt.Println(WordCount("na na na na na na na na Batman"))
	fmt.Println(WordCount("I ate a donut then I ate another donut"))
}
