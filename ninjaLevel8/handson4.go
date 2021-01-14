package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort ints and strings
	a := []int{10, 4, 6, 21, 31, 12, 14, 19}
	b := []string{`hel`, `jew`, `bar`, `foo`, `sick`}

	sort.Ints(a)
	fmt.Println(a)

	sort.Strings(b)
	fmt.Println(b)

}

