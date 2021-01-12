package main

import (
	"fmt"
)

//Build and anonymous func
func main() {

	func(){
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}

