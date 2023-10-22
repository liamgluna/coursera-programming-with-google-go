package main

import "fmt"

/*
A race condition in programming is a situation where the outcome of a program depends
on the timing of events involving multiple threads or processes,
leading to unpredictable and potentially undesirable results.

Race conditions occur when multiple threads or processes access shared resources concurrently,
and the final outcome becomes unpredictable due to the timing of their execution.
This can lead to data corruption, deadlocks, or inconsistent program states.
Race conditions happen when synchronization mechanisms aren't used to control access to shared resources,
allowing multiple threads or processes to interfere with each other's operations.
*/

func main() {
	count := 0
	n := 1000

	for i := 0; i < n; i++ {
		go func() {
			count++
		}()
	}
	fmt.Println("count:", count)
}
