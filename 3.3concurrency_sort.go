// concurrency_bubblesort
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func bubblesort(unsorted []int, wg *sync.WaitGroup) {
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
	wg.Done()
}

func swap(unsorted []int, y int) {
	unsorted[y], unsorted[y+1] = unsorted[y+1], unsorted[y]
}

func userprompt() []int {
	fmt.Print("Please input the sequence of integer numbers,\nseparated by space, like 1 89 54 5 : ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\r\n")
	input = strings.ReplaceAll(input, "  ", " ")
	inputstring := strings.Split(input, " ")
	unsorted := make([]int, 0)
	for _, y := range inputstring {
		add, _ := strconv.Atoi(y)
		unsorted = append(unsorted, add)
	}
	return unsorted
}

func main() {
	unsorted := userprompt()
	fmt.Println(unsorted)
	val := int(len(unsorted) / 4)
	volume1, volume2, volume3, volume4 := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	volume1, volume2, volume3, volume4 = unsorted[0:val], unsorted[val:val*2], unsorted[val*2:val*3], unsorted[val*3:len(unsorted)]
	fmt.Println("Array 1: ", volume1)
	fmt.Println("Array 2: ", volume2)
	fmt.Println("Array 3: ", volume3)
	fmt.Println("Array 4: ", volume4)
	var wg sync.WaitGroup
	wg.Add(4)
	go bubblesort(volume4, &wg)
	go bubblesort(volume3, &wg)
	go bubblesort(volume2, &wg)
	go bubblesort(volume1, &wg)
	wg.Wait()
	sorted := volume1
	for _, i := range volume2 {
		sorted = append(sorted, i)
	}
	for _, i := range volume3 {
		sorted = append(sorted, i)
	}
	for _, i := range volume4 {
		sorted = append(sorted, i)
	}
	wg.Add(1)
	go bubblesort(sorted, &wg)
	wg.Wait()
	fmt.Println("Sorted array: ", sorted)
}
