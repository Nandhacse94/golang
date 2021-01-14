package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	a := user{
		Name: "James",
		Age:  30,
	}
	b := user{
		Name: "Bond",
		Age:  25,
	}

	c := user{
		Name: "Alex",
		Age:  20,
	}

	d := []user{a, b, c}

	fmt.Println(d)

	sort.Sort(ByAge(d))

	fmt.Println(d)

}

