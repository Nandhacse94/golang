package main

import (
	"fmt"
)

func main() {
	//ok idiom
	c := make(chan int)

	go func() { c <- 42 }()

	v, ok := <-c
	fmt.Println(v, "\t", ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, "\t", ok)

}

