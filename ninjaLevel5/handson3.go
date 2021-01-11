package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
	
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	a := vehicle {
		doors: 4,
		color: "blue",
	}
	
	b := truck {
		vehicle: vehicle {
		doors: 2,
		color: "white",
		},
		fourwheel: true,
	}
	
	c := sedan {
		vehicle: vehicle {
		doors: 4,
		color: "black",
		},
		luxury: true,
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

