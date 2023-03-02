// bubblesort
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Function implements bubble sort algorithm
func bubblesort(unsorted []int) {
	for i := 0; i < len(unsorted)-1; i++ {
		flag := 0
		for y := 0; y < len(unsorted)-1-i; y++ {
			if unsorted[y] > unsorted[y+1] {
				flag = 1
				swap(unsorted, y)
			}
		}
		if flag == 0 {
			break
		}
	}
}

//Swapping places of neighboring values
func swap(unsorted []int, y int) {
	unsorted[y], unsorted[y+1] = unsorted[y+1], unsorted[y]
}

func main() {
	fmt.Println("Please input the sequence of integer numbers, separated by space, like 1 89 54 5 : ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	input = strings.Trim(input, "\r\n")
	input = strings.ReplaceAll(input, "  ", " ")
	inputstring := strings.Split(input, " ")
	unsorted := make([]int, 0)
	for _, y := range inputstring {
		add, _ := strconv.Atoi(y)
		unsorted = append(unsorted, add)
	}

	bubblesort(unsorted)
	fmt.Println(unsorted)
}
