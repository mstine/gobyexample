// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import "fmt"

func main() {
	// START OMIT
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// END OMIT
}
