package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	fmt.Println(&p)
	p.first = "Nandha"  // one way
	(*p).last = "kumar" //another way
	(*p).age = 26
}

func main() {
	//create a person struct
	//create a func called "changeMe" and person as parameter
	//change the value of person fields.
	//create a value of type person and print
	//call changeMe

	f := person{
		first: "karthi",
		last:  "keyan",
		age:   21,
	}
	changeMe(&f)
	fmt.Println(f)
}


