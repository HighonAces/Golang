package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("/Users/srujanreddy/Documents/Coding Projects/Golang/go_book/ch3/basename1.go"))
}
