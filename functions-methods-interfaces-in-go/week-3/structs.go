package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	animalMap := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		var animal, req string
		fmt.Print("> ")
		fmt.Scan(&animal, &req)

		animal = strings.ToLower(animal)
		req = strings.ToLower(req)

		if req == "eat" {
			fmt.Println(animalMap[animal].Eat())
		} else if req == "move" {
			fmt.Println(animalMap[animal].Speak())
		} else if req == "speak" {
			fmt.Println(animalMap[animal].Move())
		} else {
			fmt.Println("Invalid input")
		}
	}
}
