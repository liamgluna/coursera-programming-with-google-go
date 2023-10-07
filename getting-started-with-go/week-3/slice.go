package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 0, 3)

	for {
		var input string
		fmt.Print("Enter an integer (x to exit): ")
		fmt.Scan(&input)

		if input == "X" || input == "x" {
			break
		}

		var num int
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		intSlice = append(intSlice, num)

		sort.Ints(intSlice)
		fmt.Println(intSlice)
	}

}
