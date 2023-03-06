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
	fmt.Println(name, "eats: ", p.food)
}

func (p animal) move(name string) {
	fmt.Println(name, "moves: ", p.loco)
}

func (p animal) speak(name string) {
	fmt.Println(name, "speaks: ", p.sound)
}

// Func userprompt lets to input just command words in one line with space
func userprompt() (anim, animinfo string) {
	fmt.Println("Input animal(cow, bird, snake) and function (eat, move, speak)> ")
	animinfo, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	animinfo = strings.Trim(animinfo, "\r\n")
	inputstring := strings.Split(animinfo, " ")
	anim = inputstring[0]
	return
}

func main() {
	var cow = animal{food: "grass", loco: "walk", sound: "moo"}
	var bird = animal{food: "worms", loco: "fly", sound: "peep"}
	var snake = animal{food: "mice", loco: "slither", sound: "hsss"}
	var anim, animinfo string
	for {
		anim, animinfo = userprompt()

		switch animinfo {
		case "cow eat":
			cow.eat(anim)
		case "cow move":
			cow.move(anim)
		case "cow speak":
			cow.speak(anim)
		case "bir deat":
			bird.eat(anim)
		case "bird move":
			bird.move(anim)
		case "bird speak":
			bird.speak(anim)
		case "snake eat":
			snake.eat(anim)
		case "snake move":
			snake.move(anim)
		case "snake speak":
			snake.speak(anim)
		}
	}
}
