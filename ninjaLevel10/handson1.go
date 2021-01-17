package main

import (
	"fmt"
)

func main() {
	//create a working buffered channel
	a := make(chan int, 1) //1 means 1 buffer

	//to start a separate go routine, so channel works.
	go func() {
		a <- 42
	}()

	fmt.Println(<-a)
}

