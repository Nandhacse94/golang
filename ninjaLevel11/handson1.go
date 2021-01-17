package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	//Instead of using blank identifier, make sure code is checking and handling error.
	p := person{
		First:   "Nandha",
		Last:    "Kumar",
		Sayings: []string{"Hello", "I completed B.E.CSE", "Im a SDE"},
	}

	bs, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error Caught: ", err)
		return
	}

	fmt.Println("The json string is, ", string(bs))

}
