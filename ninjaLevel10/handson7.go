package main

import (
	"fmt"
)

func main() {
	//puts 100 number on a channel and prints them off of a channel

	c := make(chan int, 10)

	go foo(c)

	for i := 1; i <= 100; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("about to exit")
}

func foo(c chan<- int) {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 1; j <= 10; j++ {
				c <- i*10 + j
			}
		}(i)
	}
}

