package main

import "fmt"

var x int
var y string
var z bool

func main(){

	//DEFAULT values by COMPILER, called as ZERO VALUES.
	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(z);
	
	//User defined VALUES.
	x = 10;
	y = "Hello World";
	z = true
	
	//Printing the VALUES.
	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(z);

}

