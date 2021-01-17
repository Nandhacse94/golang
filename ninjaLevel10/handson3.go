package main

import (
	"fmt"
)

func main() {
	//create func to send the values through channels and return the receive only channel.
	//print the values from receive channel.
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {

	for v := range c {
		fmt.Println(v)
	}
}

