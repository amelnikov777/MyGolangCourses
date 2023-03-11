package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	eat()
	move()
	speak()
}
type fields struct {
	atype string
	food  string
	loco  string
	sound string
}

func (p fields) eat() {
	fmt.Print("is a ", p.atype, ". It eats ", p.food, ".\n")
}

func (p fields) move() {
	fmt.Print("is a ", p.atype, ". It moves by ", p.loco, ".\n")
}

func (p fields) speak() {
	fmt.Print("is a ", p.atype, ". It speaks ", p.sound, ".\n")
}

//Func userprompt lets to input just command words in one line with space
func userprompt(err string) (command string, name string, info string) {
	fmt.Println(err)
	command, name, info = "", "", ""
	fmt.Println("Input format:")
	fmt.Println("1. newanimal, name, type (cow, bird, snake)")
	fmt.Println("2. query, name, function (eat, move, speak)")
	fmt.Print("> ")
	inputline, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputline = strings.Trim(inputline, "\r\n")
	inputline = strings.ReplaceAll(inputline, "  ", " ")
	fmt.Println(inputline)
	inputstring := strings.Split(inputline, " ")
	if len(inputstring) != 3 {
		userprompt("\nError! Wrong format.\n")
	}
	command, name, info = inputstring[0], inputstring[1], inputstring[2]
	return
}

//Func newanimal fills properties of animal types into the map
func newanimal(info string) fields {
	switch info {
	case "cow":
		return fields{"cow", "grass", "walk", "moo"}
	case "bird":
		return fields{"bird", "worms", "fly", "peep"}
	case "snake":
		return fields{"snake", "mice", "slither", "hsss"}
	default:
		userprompt("\nError! Animal type is wrong.\n")
	}
	return fields{"", "", "", ""}
}

func main() {
	animals := make(map[string]fields)
	var command, name, info string
	for {
		err := ""
		command, name, info = userprompt(err)
		if command == "newanimal" {
			for i, _ := range animals {
				if i == name {
					userprompt("\nError! This name is already exist.\n")
				}
			}
			animals[name] = newanimal(info)
		}
		if command == "query" {
			var a Animal = fields{animals[name].atype, animals[name].food, animals[name].loco, animals[name].sound}
			fmt.Print("----------------------------------")
			fmt.Print("\n", name, " ")
			switch info {
			case "eat":
				a.eat()
			case "move":
				a.move()
			case "speak":
				a.speak()
			}
			fmt.Print("----------------------------------")
		}
		if command != "newanimal" && command != "query" {
			err = "\nError! Wrong command.\n"
		}
		fmt.Println("\nThe database contains:")
		for i, ii := range animals {
			fmt.Print(i, ii, "\n")
		}
	}
}
