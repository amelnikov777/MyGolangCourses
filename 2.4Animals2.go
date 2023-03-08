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

type f struct {
	food  string
	loco  string
	sound string
}

func (p f) eat() {
	fmt.Println("eats: ", p.food)
}

func (p f) move() {
	fmt.Println("locomotion: ", p.loco)
}

func (p f) speak() {
	fmt.Println("speaks: ", p.sound)
}

//Func userprompt lets to input just command words in one line with space
func userprompt(err string) (comand string, name string, info string) {
	fmt.Println(err)
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
		userprompt("Error! Wrong format.")
	}
	comand, name, info = inputstring[0], inputstring[1], inputstring[2]
	return
}

func newanimal(info string) fields {
	switch info {
	case "cow":
		return fields{"cow", "grass", "walk", "moo"}
	case "bird":
		return fields{"bird", "worms", "fly", "peep"}
	case "snake":
		return fields{"snake", "mice", "slither", "hsss"}
	default:
		userprompt("Error! Animal type is wrong.")
	}
	return fields{"", "", "", ""}
}

func main() {
	animals := make(map[string]fields)
	var comand, name, info string
	err := ""
	for {
		comand, name, info = userprompt(err)
		if comand == "newanimal" {
			for i, _ := range animals {
				if i == name {
					userprompt("Error! This name is already exist.")
				}
			}
			animals[name] = newanimal(info)
		}
		if comand == "query" {
			var a Animal = f{animals[name].food, animals[name].loco, animals[name].sound}
			switch info {
			case "eat":
				a.eat()
			case "move":
				a.move()
			case "speak":
				a.speak()
			}
		}
		fmt.Println("Current database contents:")
		for i, ii := range animals {
			fmt.Print(i, ii, "\n")
		}
	}
}
