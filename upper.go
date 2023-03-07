package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithProg := os.Args[1:]
	var check = false
	for _, a := range argsWithProg {
		if strings.ToLower(a) == "lower" {
			check = true;
		}
	}

	if check == false {
		fmt.Println(strings.ToUpper(strings.Join(argsWithProg, " ")))
	} else {
		fmt.Println(strings.ToLower(strings.Join(argsWithProg, " ")))
	}
}
