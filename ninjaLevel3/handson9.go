package main

import (
	"fmt"
)

func main() {
	//switch statement with switch expression sepcified as TYPE string and IDENTIFIER favSport.

	favSport := "FootBall"

	switch favSport {
	case "Cricket":
		fmt.Println("favSport is Cricket")
	case "Hockey":
		fmt.Println("favSport is Hockey")
	case "FootBall":
		fmt.Println("favSport is Football")
	default:
		fmt.Println("This is default case when above cases fails")
	}
}
