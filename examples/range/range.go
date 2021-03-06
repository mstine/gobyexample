// _range_ iterates over of elements in a variety of
// data structures. Let's see how to use `range` with some
// of the data structures we've already learned.

package main

import "fmt"

func main() {
	// START OMIT
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// END OMIT

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
