package main

import (
	"fmt"
)

func main() {
	a := struct {
		numbers []int
		city    map[string]int
	}{
		numbers: []int{4, 3, 2, 1},

		city: map[string]int{
			"cse": 4,
			"sde": 4,
		},
	}

	fmt.Println(a)
	
	for k,v := range a.city{
		fmt.Println(k, v);
	}

}
