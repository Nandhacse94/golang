package main

import (
	"fmt"
)

func main() {
	//program with now switch expression specified.

	switch {
	case 1 == 3:
		fmt.Println("it wont print 1 == 3")
	case true:
		fmt.Println("no switch expression, then true will be passed into case")
	case 1 < 3:
		fmt.Println("it will print as 1 < 3")
	default:
		fmt.Println("This is default case when above cases fails")
	}
}

