package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithProg := os.Args[1:]
	fmt.Println(strings.ToUpper(strings.Join(argsWithProg, " ")))
}
