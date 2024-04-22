package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] + strconv.Itoa(len(os.Args))
		sep = " "
	}
	fmt.Println(s)
}
