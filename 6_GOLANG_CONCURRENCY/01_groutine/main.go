package main

import (
	"fmt"
	"time"
)

func main() {
	//doing multiple tasks in a way that thay can overlap in time, instead of doing them one by one

	/*
		concurrency: doing multiple tasks at same time but not necessarily at the same time
		parallelism: doing multiple tasks at the same time
	*/

	// golang code can be concurrect without being parallel, but it can be parallel without being concurrent

	/*
			when not to use concurrency:
		- when you have a simple task that can be done in a single thread
		- when you have a task that is not CPU bound
		- when you have a task that is not I/O bound
		- when you have a task that is not network bound
		- when you have a task that is not memory bound
	*/

	/*
		what is a goroutine: a lightweight thread managed by the Go runtime. It is a function that is capable of running concurrently with other functions. Goroutines are cheaper than threads and can be created in large numbers without significant overhead. They are multiplexed onto a smaller number of OS threads, which allows for efficient use of system resources.
	*/

	// in simple words goroutine is an concurrent execution of a function

	/*
			rules for goroutines:
		- a goroutine is created using the go keyword followed by a function call
		- a goroutine runs concurrently with the calling function
		- a goroutine can communicate with other goroutines using channels
		- a goroutine can be synchronized using wait groups
		- a goroutine can be canceled using context
	*/

	start := time.Now()

	go func() {
		time.Sleep(300 * time.Millisecond)

		fmt.Println("goroutine A: finished simulated Api at:",time.Since(start))
	}()

	go func(){
		time.Sleep(150 * time.Millisecond)
		fmt.Println("goroutine B: finished simulated Api at:",time.Since(start))
	}()

	fmt.Println("main: started two go routing started at:",time.Since(start))

	fmt.Println("main: doing step 1:",time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main: doing step 2:",time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main: doing step 3:",time.Since(start))
	time.Sleep(100 * time.Millisecond)

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main: exiting at:",time.Since(start))
	}