package main

import "fmt"

// gofmt -w main.goでフォーマットを実行
func main() {
	/*
		x := 1 + 1
		fmt.Println(x)
	*/

	// ビットシフト
	// 1を0回シフトする
	fmt.Println(1 << 0) // 0001 0001
	// 1を1回シフトする
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000

}
