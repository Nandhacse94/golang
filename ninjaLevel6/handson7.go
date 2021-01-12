package main

import (
	"fmt"
)

//Assign a func to a variable and call that func
func main() {

	a := func(n int) {
		for i := 1; i <= n; i++ {
			fmt.Println(i)
		}
	}

	a(100)
	fmt.Println("done")
}
