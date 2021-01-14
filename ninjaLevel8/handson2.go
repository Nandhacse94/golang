package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	a := person{"Nandha",
		"Kumar",
		26,
	}
	b := person{
		First: "James",
		Last:  "Bond",
		Age:   35,
	}
	c := person{
		First: "Go",
		Last:  "lang",
		Age:   20,
	}

	t := []person{a, b, c}
	r, e := json.Marshal(t)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(r))
	
	//Unmarshal the json to struct
	var per []person
	err := json.Unmarshal([]byte(r),&per)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(per)
}
