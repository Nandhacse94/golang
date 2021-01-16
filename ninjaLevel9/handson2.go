package main

import (
	"fmt"
)


type person struct {
	Name string
	Age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("The Speaker is ", p.Name, " and  ", p.Age, " years Old")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	//method sets and pass a pointer to the method to get it worked
	p := person{
		Name: "Nandha",
		Age:  26,
	}
	saySomething(&p)
}

