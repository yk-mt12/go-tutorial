package main

import (
	"fmt"
)

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	n := 10
	fmt.Println(isEven(n))
}
