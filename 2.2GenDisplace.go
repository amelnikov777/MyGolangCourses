// GenDisplace
package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	dist := func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
	return dist
}

func main() {
	var a, vo, so, t float64
	var Y string = "Y"
	for Y == "Y" || Y == "y" {
		fmt.Print("\nPlease enter values")
		fmt.Print("\nAcceleration : ")
		fmt.Scanln(&a)
		fmt.Print(a)
		fmt.Print("\nInitial velocity : ")
		fmt.Scanln(&vo)
		fmt.Print(vo)
		fmt.Print("\nInitial displacement : ")
		fmt.Scanln(&so)
		fmt.Print(so)
		fmt.Print("\nEnter on the way time to calculate displacement : ")
		fmt.Scanln(&t)
		fmt.Print(t)

		fn := GenDisplaceFn(a, vo, so)
		fmt.Print("\nDisplacenement is : ", fn(t), "\n")

		fmt.Print("\nPress Y to repeat : ")
		fmt.Scanln(&Y)
	}
	fmt.Print("\n The End \n")
}
