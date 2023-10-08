package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	intSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		var input string
		fmt.Print("Enter an integer (x to stop): ")
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

	}
	BubbleSort(intSlice)
	fmt.Println(intSlice)

}

func BubbleSort(intSlice []int) {
	for i := 1; i < len(intSlice); i++ {
		for j := 0; j < len(intSlice)-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}
		}
	}
}

func Swap(intSlice []int, idx int) {
	temp := intSlice[idx]
	intSlice[idx] = intSlice[idx+1]
	intSlice[idx+1] = temp

}