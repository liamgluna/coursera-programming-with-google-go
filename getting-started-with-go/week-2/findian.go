package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	fmt.Print("Enter a string: ")

	// Using bufio for edge case of input with multiple spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	word = scanner.Text()

	word = strings.ToLower(word)
	word = strings.TrimSpace(word)


	if strings.HasPrefix(word, "i") && strings.HasSuffix(word, "n") && strings.Contains(word, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
