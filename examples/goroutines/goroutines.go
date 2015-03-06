// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"

// START_F OMIT
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// END_F OMIT

func main() {
	done := make(chan string)
	// START_MAIN OMIT
	// Synchronous
	f("direct")

	// Asynchronous
	go f("goroutine")

	// Anonymous and Asynchronous
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// END_MAIN OMIT
	<-done
}

// END OMIT
