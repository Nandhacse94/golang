package main

import (
	"fmt"
)

type numbers int
var x numbers;

func main() {
	fmt.Println(x);
	fmt.Printf("%T\n", x);

	x = 42;
	
	fmt.Println(x);
	fmt.Printf("%T\n", x);

}

