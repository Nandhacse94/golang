package main

import (
	"fmt"
)

func main() {
	//variadic argument type of int and pass as argument
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := foo(a...) //unfurl

	fmt.Println(b)

	//variadic argument type of int and pass as argument
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := bar(c) //raw slice

	fmt.Println(d)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
