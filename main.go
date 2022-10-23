package main

import "fmt"

/*
一番最初に実行される特別な関数
設定ファイルを記載する
*/

func init() {
	fmt.Println("Init!")
}

func buzz() {
	fmt.Println("Buzz")
}

func main() {
    fmt.Println("hello, world")
		buzz()
}