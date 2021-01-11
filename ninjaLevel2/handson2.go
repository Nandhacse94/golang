package main

import (
	"fmt"
)

func main() {
	//use operators as expression and store the result into a variable.
	a := 10;
	b := 20;
	
	c := (a == b);
	d := (a != b);
	e := (a < b);
	f := (a > b);
	g := (a >= b);
	h := (a <= b);
	
	fmt.Println(c);
	fmt.Println(d);
	fmt.Println(e);
	fmt.Println(f);
	fmt.Println(g);
	fmt.Println(h);
	
}
