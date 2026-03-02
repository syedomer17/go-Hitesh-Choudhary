package main

import (
	"fmt"
	"go-modules/internal/greet"
)

func main() {
	msg1 := greet.Hello("omer")

	fmt.Println(msg1)
}