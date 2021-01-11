package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	ice_flavour []string
}

func main() {

	p1 := person{
		first:       "nandha",
		last:        "kumar",
		ice_flavour: []string{ "vanilla", "strawberry", "mango",},
	}

	p2 := person{
		first:       "siva",
		last:        "kumar",
		ice_flavour: []string{ "chocho", "blue", "fruit",},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	
	for k,v := range p1.ice_flavour {
		fmt.Println(k, v)
	}

}

