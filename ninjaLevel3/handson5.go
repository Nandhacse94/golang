package main

import (
	"fmt"
)

func main() {
	//print remainder from the number between 10 and 10000 when its divided by 4.

	for i := 10; i <= 10000; i++ {

		fmt.Printf("number %v and remainder after divided by four = %v\n", i, i%4)
	}
}

