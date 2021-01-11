package main

import (
	"fmt"
)

func main() {
	//print the years you have been alive.

	born := 1994
	present := 2021

	for {
		if born > present {
			break
		}
		fmt.Println(born)
		born++
	}
}

