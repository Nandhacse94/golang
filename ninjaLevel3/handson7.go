package main

import (
	"fmt"
)

func main() {
	//program with if-statemnt with else if and else.
	
	str := "James Bond";
	
	if str == "Bond"{
		fmt.Println("Bond")
	}else if str == "James"{
		fmt.Println("James")
	}else if str == "James Bond"{
		fmt.Println("James Bond")
	}else {
		fmt.Println("there is no Bond or James or James bond");
	}
}

