package main

import (
	"fmt"
)

func main() {
	//create an SLICE which hold 5 values with type of INT

	a := []int{10,11,12,13,14,15,16,17,18,19,20}

	fmt.Println(a)

	//print the values using range
	for i, v := range a {
		fmt.Println(i, v)
	}
	//type of the array
	fmt.Printf("%T\n", a)
}

