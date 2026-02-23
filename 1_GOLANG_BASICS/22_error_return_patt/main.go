package main

import (
	"fmt"
	"strconv"
)

func main() {
	// go don't use exception for normal failurs
	// function returns errors as normal return values

	// val, err := something()
	// if err != nil {
	// 	 handle error
	// }

	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func run() error {

	input := "5"

	level, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println("selected level", level)
	return nil
}

func parseLevel(s string) (int, error) {

	// nil error => success
	// non nil error => failure

	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("The level must be between 1 and 5")
	}

	return n, nil
}