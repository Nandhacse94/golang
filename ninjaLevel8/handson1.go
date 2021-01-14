package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	//define a struct and create 3 values and convert to json and print
	a := person{"Nandha",
		"Kumar",
		26,
	}
	b := person{
		First: "James",
		Last:  "Bond",
		Age:   35,
	}
	c := person{
		First: "Go",
		Last:  "lang",
		Age:   20,
	}

	t := []person{a, b, c}
	r, e := json.Marshal(t)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(r))

}
