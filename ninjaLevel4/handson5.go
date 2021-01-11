package main

import (
	"fmt"
)

func main() {
	//create a SLICE
	a := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println(a)

	//append and remove some values.
	a = append( a[:4], a[6:]... )

	fmt.Println(a)

}
