package main

import (
	"fmt"
)

func main() {
	//conditions over println functions and check what it prints.

	fmt.Printf("true && true = %v\n", true && true)
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("true || true = %v\n", true || true)
	fmt.Printf("true || false = %v\n", true || false)
	fmt.Printf("!true = %v\n", !true)
}
