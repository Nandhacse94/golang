package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	//Instead of using blank identifier, make sure code is checking and handling error.
	p := person{
		First:   "Nandha",
		Last:    "Kumar",
		Sayings: []string{"Hello", "I completed B.E.CSE", "Im a SDE"},
	}

	bs, err := toJSON(p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The json string is, ", string(bs))

}


func toJSON(a interface{}) ([]byte,error){
	bs,err := json.Marshal(a)
	
	if err != nil{
		return []byte{},fmt.Errorf("Conversion to JSON error %v",err)
	}
	fmt.Println("JSON is, ",string(bs))

	return bs,nil
}
