package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("case 1: success")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\ncase 2: failure")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {

	fmt.Println("start: resource acquired")

	// defer is a powerfull funcation thet allows you to schedule a function call to be run after the function completes. It is often used for cleanup tasks, such as releasing resources or closing files.
	defer fmt.Println("cleanup: resource released")

	if !success {
		return errors.New("Something went wrong. I am returning early")
	}

	fmt.Println("work: doing somthing.")
	fmt.Println("end: work done")

	return nil
}