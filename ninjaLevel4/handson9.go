package main

import (
	"fmt"
)

func main() {
	//create a MAP with type STRING (last_name) and value STRING (fav_thing)
	
	m := map[string][]string {
	"nandha" : []string{"icecream","cake","shawarma"},
	"karthi" : []string{"grill","juice","chilly"},
	"siva" : []string{"choco","milk","popcorn"},
	}
	
	// add additional record to the map and check prints.
	m ["moorthi"] = []string{`groundnut`,`sprite`,`tea`}
	
	for k,v := range m {
		fmt.Println(k)
		for _, v1 := range v {
			fmt.Printf("\t\t%v\n",v1);
		}
	}

}

