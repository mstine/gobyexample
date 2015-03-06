// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"

// START OMIT
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Synchronous
	f("direct")

	// Asynchronous
	go f("goroutine")

	// Anonymous and Asynchronous
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println("done")
}

// END OMIT
