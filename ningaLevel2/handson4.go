package main

import (
	"fmt"
)

func main() {

	//assign a int to a variable.
	a := 8;
	
	//print that in dec, bin, hex
	fmt.Printf("%d\t%b\t%#X\n", a, a, a);
	
	//shift that bits over one position to the left and assign it to a variable.
	b := a << 1;
	
	//print that  new variable in dec, bin, hex
	fmt.Printf("%d\t%b\t%#X\n", b, b, b);	
}
