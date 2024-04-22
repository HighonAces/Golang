package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += strconv.Itoa(i) + sep + arg
		sep = " "
	}
	fmt.Println(s)
}
