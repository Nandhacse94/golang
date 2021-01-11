package main

import (
	"fmt"
)

func main() {

	//create a string SLICE using MAKE
	a := make([]string, 50, 50)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = []string{`hello`, `world`, `usa`, `uk`, `america`, `japan`}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	//print without range
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

}

