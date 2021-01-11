package main

import (
	"fmt"
)

func main() {
	//create an SLICE which hold 11 values with type of INT
	a := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println(a)

	//slicing the values and create new slices.
	b := a[:5]
	fmt.Println(b)

	c := a[6:]
	fmt.Println(c)

	d := a[3:8]
	fmt.Println(d)

	e := a[5:10]
	fmt.Println(e)
}

