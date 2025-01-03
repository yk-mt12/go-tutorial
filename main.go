package main

import (
	"fmt"
	"flag"
	"strings"
)

var msg = flag.String("msg", "hoge", "Flag to print a message")
var n int

func init() {
	flag.IntVar(&n, "n", 1, "回数")
}
func main() {
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
}
