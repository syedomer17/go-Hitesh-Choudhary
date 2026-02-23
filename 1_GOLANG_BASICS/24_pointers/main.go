package main

import "fmt"

func main() {
	// store the memory addess of any value

	// &x -> is a pointer of x (makes a pointer)
	// *p -> deref (go to that address and read/write )

	score := 10
	fmt.Println("before:",score)

	addScore(&score)
}

func addScore(score *int) {
	*score = *score + 10 
	fmt.Println("after:",*score)
}