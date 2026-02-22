package main

import "fmt"

func main() {
	views := []int{100, 200, 300, 400, 500}

	// for range
	total := 0

	for index, value := range views {
		fmt.Println("day", index,"views", value)

		total = total + value
	}
	fmt.Println("Total views", total)
}