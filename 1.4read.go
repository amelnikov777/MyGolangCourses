//read from file
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type readst struct {
		Iname, Fname string
	}
	var read readst
	readstslice := make([]readst, 0)
	var filename string

	fmt.Print("Please input the name of file to read: ")
	fmt.Scan(&filename)
	fmt.Println(" ", filename)
	f, err := os.Open(filename)
	if err == nil {
		readdata := bufio.NewScanner(f)

		for readdata.Scan() {
			read.Iname, read.Fname, _ = strings.Cut(readdata.Text(), " ")
			readstslice = append(readstslice, read)
		}
		for i, ii := range readstslice {
			fmt.Println(i, ii)
		}
	} else {
		fmt.Println("File open error! Please check the filename.")
	}
}
