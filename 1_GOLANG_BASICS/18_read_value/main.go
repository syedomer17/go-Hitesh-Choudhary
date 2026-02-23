package main

import (
	"fmt"
)

func main() {
	points := map[string]int{
		"a": 10,
		"b": 20,
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"]) // this will return 0 because c is not exist in the map

	val, okb := points["b"] // this will return the value of b and true because b is exist in the map
	fmt.Println("b", val, okb)

	val, okc := points["c"]
	fmt.Println("c", val, okc) // this will return 0 and false because c is not exist in the map

	if val, ok := points["b"]; ok {
		fmt.Println("b exists with value:", val)
	}else{
		fmt.Println("b does not exist in the map")
	}

	prices := map[string]int{
		"xyz": 5000,
		"def": 1800,
	}

	total := 0

	for item, price := range prices {
		fmt.Printf("Item: %s, Price: %d\n", item, price)
		total = total + price
	}
	fmt.Println("Total:", total)

	for items := range prices {
		fmt.Println(items)
	}
}