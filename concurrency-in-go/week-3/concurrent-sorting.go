/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	// Read the integers from the user
	integers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Enter an integer: ")
		fmt.Scan(&integers[i])
	}

	// Divide the array into four subarrays
	subarraySize := len(integers) / 4
	subarrays := make([][]int, 4)

	for i := 0; i < 4; i++ {
		start := i * subarraySize
		end := (i + 1) * subarraySize
		if i == 3 {
			end = len(integers)
		}
		subarrays[i] = integers[start:end]
	}

	// Create a WaitGroup to wait for the goroutines to finish
	var wg sync.WaitGroup

	// Sort each subarray in a separate goroutine
	fmt.Println("Subarrays:")
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(subarray []int) {
			defer wg.Done()
			fmt.Println(subarray)
			sort.Ints(subarray)
		}(subarrays[i])
	}

	// Wait for all sorting goroutines to finish
	wg.Wait()

	// Merge the sorted subarrays
	merged := merge(subarrays)

	// Print the entire sorted list
	fmt.Println("Sorted List:", merged)
}

func merge(subarrays [][]int) []int {
	result := make([]int, 0)
	for _, subarray := range subarrays {
		result = mergeTwoSortedArrays(result, subarray)
	}
	return result
}

func mergeTwoSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, 0)
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}
