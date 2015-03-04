// _Functions_ are central in Go. We'll learn about
// functions with a few different examples.

// START OMIT
package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	result := plus(1, 2)
	fmt.Println("1 + 2 =", res)
}

// END OMIT
