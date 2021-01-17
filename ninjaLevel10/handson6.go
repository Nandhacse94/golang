package main

import (
	"fmt"
)

func main() {
	//puts 100 number on a channel and prints them off of a channel

	c := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

