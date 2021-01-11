package main

import (
	"fmt"
)

func main() {
	//create a SLICE
	a := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	//append
	a = append(a, 21)

	fmt.Println(a)

	a = append(a, 22, 23, 24, 25)

	fmt.Println(a)

	b := []int{26, 27, 28, 29, 30}

	fmt.Println(b)

	//append other slice to slice identifier 'a'
	a = append(a, b[0:]...)

	fmt.Println(a)

}

