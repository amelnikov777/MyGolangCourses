package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food  string
	loco  string
	sound string
}

func (p animal) eat(name string) {
	fmt.Println(name, " eats: ", p.food)
}

func (p animal) move(name string) {
	fmt.Println(name, " locomotion: ", p.loco)
}

func (p animal) speak(name string) {
	fmt.Println(name, " speak: ", p.sound)
}

//Func userprompt lets to input just correct command words in one line with space
func userprompt() (anim string, info string) {
	var commands = []string{"cow eat", "cow move", "cow speak", "bird eat", "bird move", "bird speak", "snake eat", "snake move", "snake speak"}
	t := false
	fmt.Println("Input animal(cow, bird, snake) and function (eat, move, speak)> ")
	inputline, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputline = strings.Trim(inputline, "\r\n")
	inputline = strings.ReplaceAll(inputline, "  ", " ")
	for _, i := range commands {
		if i == inputline {
			t = true
			break
		}
	}
	if t == false {
		fmt.Println("Invalid command. Try again.\n")
		return
		userprompt()
	}
	inputstring := strings.Split(inputline, " ")
	anim, info = inputstring[0], inputstring[1]
	return
}

func main() {
	var animals = map[string]animal{
		"cow":   animal{"grass", "walk", "moo"},
		"bird":  animal{"worms", "fly", "peep"},
		"snake": animal{"mice", "slither", "hsss"},
	}

	var anim, info string
	for {
		anim, info = userprompt()

		switch info {
		case "eat":
			animals[anim].eat(anim)
		case "move":
			animals[anim].move(anim)
		case "speak":
			animals[anim].speak(anim)
		}
	}
}
