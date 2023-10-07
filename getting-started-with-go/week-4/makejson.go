package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the
name and address to the map using the keys "name" and "address", respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

func main() {
	person := make(map[string]string)

	var name string
	var address string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Enter address: ")
	scanner.Scan()
	address = scanner.Text()

	person["name"] = name
	person["address"] = address

	b, err := json.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

}
