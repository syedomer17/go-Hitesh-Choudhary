package main 

import (
	"fmt"
)

// In go arrays are haiving fixed size and also the data type of the elements in an array should be the same data type.

func main(){
	// fixed and can not grow 
	var marks[3]int

	marks[0] = 10 
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	res := [5]int{1,2,3,4,5}
	fmt.Println(len(res))
}