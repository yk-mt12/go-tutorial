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
	// flagで指定された値分、msgを出力
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))

	// flag.Args関数を使用し、フラグ以外の引数を取得
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(args)
  }
}
