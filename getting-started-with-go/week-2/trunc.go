package main

import (
	"fmt"
)

func main() {
	var num float32

	fmt.Print("Enter an floating point number: ")
	fmt.Scan(&num)

	// Print and convert to integer in order to truncate
	fmt.Printf("You entered %d\n", int(num))
}
