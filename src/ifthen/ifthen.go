package main

import (
	"fmt"
)

func main() {

	var test int = 0
	var msg string

	if test > 0 {
		msg = "Greater than zero"

	} else if test == 0 {

		msg = "Equal to zero"
	} else {
		msg = "Lesser than zero"
	}

	fmt.Println(msg)

}
