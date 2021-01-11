package main

import (
	"fmt"
)

func main() {
	//two slice
	a := []string{`james`, `bond`, `shaken`, `not stirred`}
	b := []string{`miss`, `money`, `penny`, `cool`}
	//2d slice
	c := [][]string{a, b}

	fmt.Println(c)

	//print each value using range
	for i, v := range c {
		for k, v1 := range v {
			fmt.Println(i, k, v1)
		}
		fmt.Println()
	}
}

