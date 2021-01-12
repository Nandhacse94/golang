package main

import (
	"fmt"
)

// create a struct with identifier person - fields(first,last,age)
type person struct {
	first string
	last  string
	age   int
}

//create a func with a receiver of the person struct.
func (p person) speak() {
	fmt.Println("I'm ", p.first, p.last, " and my age is ", p.age)
}

func main() {

	//create a value and call the method
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   26,
	}

	p1.speak()
}
