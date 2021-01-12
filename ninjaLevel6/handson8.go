package main

import (
	"fmt"
)

// create a func which returns func
func getfunc() func() int {
	return func() int {
		return 158
	}
}

func main() {
	//call the func and store the returned func and call it.
	a := getfunc()
	fmt.Println(a())
}

