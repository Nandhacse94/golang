package main

import (
	"fmt"
)

func main() {
	//create an ARRAY which hold 5 values with type of INT

	a := [5]int{}

	fmt.Println(a)

	//assign values to array
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	//print the values using range
	for i, v := range a {
		fmt.Println(i, v)
	}
	//type of the array
	fmt.Printf("%T\n", a)
}
