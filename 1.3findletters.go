//Find "ian"
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var UserStrNew string
	fmt.Println("Please input some string: ")
	UserStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Input string is: ", UserStr)

	if err == nil {
		UserStrNew = strings.Replace(strings.ToLower(strings.TrimSpace(UserStr)), " ", "", -1)
	} else {
		fmt.Println("Error: ", err)
	}
	fmt.Println("String for analize is: ", UserStrNew, "\n")

	if strings.HasPrefix(UserStrNew, "i") && strings.Contains(UserStrNew, "a") && strings.HasSuffix(UserStrNew, "n") {
		fmt.Println("Found!\n")
	} else {
		fmt.Println("Not Found!\n")
	}
}
