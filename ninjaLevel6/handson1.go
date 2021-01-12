package main

import (
	"fmt"
)

func main() {
	//functions with returning single and multiple values.
	a := foo()
	b, c := bar()

	fmt.Println(a)
	fmt.Println(b, c)

}

func foo() int {
	return 45
}

func bar() (int, string) {
	return 10, "Good Progression"
}

