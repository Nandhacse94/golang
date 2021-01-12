package main

import (
	"fmt"
)

func main() {
	//create a func and store into one variable
	a := func(n int) int {
		return n * 15
	}

	//pass a func as argument
	n := foo(a)
	fmt.Println(n)

}

func foo(f func(n int)int) int {
	// call the func with value
	b := f(10)
	b++
	return b
}

