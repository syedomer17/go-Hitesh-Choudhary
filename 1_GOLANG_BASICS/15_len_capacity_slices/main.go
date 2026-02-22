package main 

import (
	"fmt"
)

func main(){
	scores := make([]int,0,5) // here we are creating a slice with length 0 and capacity 5 

	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 100)
	fmt.Println("After appeding 100", scores, len(scores), cap(scores))
	scores = append(scores, 200, 3000)
	fmt.Println("After appeding 200 and 3000", scores, len(scores), cap(scores))

	// if in case we are exceding the capacity go grows the backing array and copy the existing elements to the new array and then add the new elements to the new array and also the capacity of the new array is double the capacity of the old array 

	// it just doubles the capacity of the old array and copy that and make a new array and copy the existing elements to th new array and then add the new elements to the new array

	scores = append(scores, 400, 500, 600)
	fmt.Println("After appeding 400, 500 and 600", scores, len(scores), cap(scores))

	scores = append(scores, 2)
	fmt.Println("After appeding 2", scores, len(scores), cap(scores))

	todos := []string{"task1", "task2","task3"}

	more := []string{"task4", "task5"}

	todos = append(todos, more...)
	fmt.Println(todos)
} 