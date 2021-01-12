package main

import (
	"fmt"
)

//create a func that encloses the scope of the variable.
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {

	a := foo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	b := foo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

