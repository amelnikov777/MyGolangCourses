//JSONexport .
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	person := make(map[string]string)
	for {
		fmt.Print("Please input name: ")
		UserNam, errNam := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println(UserNam)
		fmt.Print("Please input address: ")
		UserAdr, errAdr := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println(UserAdr)
		if errNam == nil && errAdr == nil {
			person["address"] = strings.Trim(UserAdr, "\r\n")
			person["name"] = strings.Trim(UserNam, "\r\n")
			break
		} else {
			fmt.Println("Error: ", errNam, errAdr)
		}
	}
	JS, err := json.Marshal(person)
	if err == nil {
		fmt.Println(string(JS))
	} else {
		fmt.Println("Error JSON translate :", JS)
	}
}
