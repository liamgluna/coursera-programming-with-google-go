package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	animalMap := make(map[string]Animal)

	for {
		var command, name, typeOrInfo string
		fmt.Print("> ")
		fmt.Scan(&command, &name, &typeOrInfo)

		command = strings.ToLower(command)
		name = strings.ToLower(name)
		typeOrInfo = strings.ToLower(typeOrInfo)

		if command == "newanimal" {
			switch typeOrInfo {
			case "cow":
				animalMap[name] = Cow{
					food:       "grass",
					locomotion: "walk",
					noise:      "moo",
				}
				fmt.Println("Created it!")
			case "bird":
				animalMap[name] = Bird{
					food:       "worms",
					locomotion: "fly",
					noise:      "peep",
				}
				fmt.Println("Created it!")
			case "snake":
				animalMap[name] = Snake{
					food:       "mice",
					locomotion: "slither",
					noise:      "hsss",
				}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal")
			}
		} else if command == "query" {
			switch typeOrInfo {
			case "eat":
				animalMap[name].Eat()
			case "move":
				animalMap[name].Move()
			case "speak":
				animalMap[name].Speak()
			default:
				fmt.Println("Invalid query")
			}
		} else {
			fmt.Println("Invalid command")
		}
	}
}
