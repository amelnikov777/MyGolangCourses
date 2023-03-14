/*
Race condition is a problem where the outcome of the program depends on the interleaving.
Because, the interleaving is non-deterministic, it's determined by the operating system.
So, a race condition is a problem that can happen as a result of all these possible
interleaving which depend from system condition, but not from code.

This program runs two copies of the same function simultaneously in different go-rutins.
Despite the fact that their initial conditions are the same, the result can be different.
And not only between copies, but also between attempts of the same copy. Because the
system allocates an undefined amount of resources to run each go-rutin's cycle each time.

Depending on the performance of your computer, you can change the value of the "pause"
variable in the loop, affecting the result of the program execution.
*/

package main

import (
	"fmt"
)

func doit(x *int) {
	for i := 0; i <= 1000000; i++ {
		*x = i
	}
	return
}

func main() {
	for x := 0; x < 5; x++ {
		fmt.Println("New game.........................")
		var g1, g2 int = 0, 0
		fmt.Printf("srart:  \t%d\t%d", g1, g2)
		fmt.Println("")
		go doit(&g1)
		go doit(&g2)

		for pause := 0; pause < 1000000; pause++ { //Change this max to calibrating program speed to your system speed
		}
		fmt.Printf("finish: \t%d\t%d", g1, g2)
		fmt.Println("")
	}
}
