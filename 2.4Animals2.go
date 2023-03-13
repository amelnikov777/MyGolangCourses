package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var command, name, info string = "", "", ""

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

func userprompt() (command, name, info string) {
	command, name, info = "", "", ""
	fmt.Println("1. newanimal, name, type (cow, bird, snake)")
	fmt.Println("2. query, name, function (eat, move, speak)")
	fmt.Print("Please enter command > ")
	inputline, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputline = strings.Trim(inputline, "\r\n")
	inputline = strings.ReplaceAll(inputline, "  ", " ")
	fmt.Println(inputline)
	inputstring := strings.Split(inputline, " ")
	if len(inputstring) == 3 {
		command, name, info = inputstring[0], inputstring[1], inputstring[2]
	}
	return
}

func checkprompt(command, name, info string, animal_slices []string) bool {
	switch command {
	case "newanimal":
		{
			if info != "cow" && info != "bird" && info != "snake" {
				fmt.Println("\nError! Wrong animal.\n")
				return true
			}
			for _, i := range animal_slices {
				if i == name {
					fmt.Println("\nError! Existed name.\n")
					return true
				}
			}
			return false
		}

	case "query":
		{
			if info != "eat" && info != "move" && info != "speak" {
				fmt.Println("\nError! Wrong query.\n")
				return true
			}
			for _, i := range animal_slices {
				if i == name {
					return false
				}
			}
			fmt.Println("\nError! Name is unknown.\n")
			return true
		}
	}
	fmt.Println("\nError! Wrong comand.\n")
	return true
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
	}
	return fields{"", "", "", ""}
}

func main() {
	animals := make(map[string]fields)
	err := true
	line := "----------------------------------"

	for { //Main loop for prompt
		for err == true { //Loop for entering and cheking the command
			command, name, info = userprompt()
			var animal_slice []string = make([]string, 0)
			for i, _ := range animals {
				animal_slice = append(animal_slice, i)
			}
			err = checkprompt(command, name, info, animal_slice)
		}

		//Command "query" process
		if command == "query" {
			var a Animal = fields{animals[name].atype, animals[name].food, animals[name].loco, animals[name].sound}
			for i, _ := range animals {
				if i == name {
					fmt.Print(line)
					fmt.Print("\n", name, " ")
					switch info {
					case "eat":
						a.eat()
					case "move":
						a.move()
					case "speak":
						a.speak()
					}
					fmt.Print(line)
					break
				}
			}
		} else {
			//Command "newanimal" process
			animals[name] = newanimal(info)
		}

		//Printing current database
		fmt.Println("\nThe database contains:")
		for i, ii := range animals {
			fmt.Println(i, ii)
		}
		fmt.Println("")
		err = true
	}
}
