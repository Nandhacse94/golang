package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	ice_flavour []string
}

func main() {

	p1 := person{
		first:       "nandha",
		last:        "kumar",
		ice_flavour: []string{"vanilla", "strawberry", "mango"},
	}

	p2 := person{
		first:       "siva",
		last:        "kumar1",
		ice_flavour: []string{"chocho", "blue", "fruit"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range p1.ice_flavour {
		fmt.Println(k, v)
	}

	//map with user defined structs
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m[p1.last])
}
