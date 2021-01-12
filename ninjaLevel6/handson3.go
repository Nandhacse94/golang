package main

import (
	"fmt"
)

func main() {
	//use defer func and show that deferred func runs after the containing func exits.

	defer foo()

	bar()

	fmt.Println("This is the last line of main func.")
}
func foo() {
	fmt.Println("Hey I'm foo func!")
}

func bar() {
	fmt.Println("Hey I'm bar func!")
}

