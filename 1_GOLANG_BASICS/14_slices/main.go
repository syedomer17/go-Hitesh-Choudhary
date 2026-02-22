package main 

import (
	"fmt"
)

// In go slices are haiving dynamic size they don't have a fixed size and also the data type of the elements in a slice should be the same data type.

func main(){
	result := []string{"apple", "banana", "orange"}
	fmt.Println(result, result[0], result[len(result) - 1])

	result[1] = "grape"
	fmt.Println(result)

	var numbs []int 

	numbs = append(numbs, 10)
	numbs = append(numbs, 20)

	fmt.Println(numbs)
}