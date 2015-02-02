package main

import (
	"fmt"
)

type Record struct {
	First, Last string
	ID int
}

type AllRecords struct {
	Entries []Record
	nextID int
}

// fmt.Println knows how to interpret the output of this String function which
// returns a string and uses Sprintf
func (r *Record) String() string{
	return fmt.Sprintf("%d, %s, %s", r.ID, r.Last, r.First)
}

func (r *AllRecords) AddRecord(first, last string) {
	n := Record{
		First: first,
		Last: last,
		ID: r.nextID,
	}
	r.Entries = append(r.Entries, n)
	r.nextID++
}

// Doak's workshop - day 1, exercises with records
func main() {
	// 1st exercise
	a := Record{
		ID: 1,
		First: "John",
		Last: "Doe",
	}
	fmt.Println(&a)

	// 2nd exercise
	var num int64
	pointer := &num
	num = 5
	fmt.Println(num)
	*pointer = 7
	fmt.Println(num)

	// 3rd exercise
	records := &AllRecords{} // created empty, slice Entries has no values,
	                         // nextID is 0
	records.AddRecord("Flash", "Gordon")
	records.AddRecord("Hawk", "Eye")
	records.AddRecord("GI", "Joe")
	for _, employee := range records.Entries {
		fmt.Println(&employee)
	}
}