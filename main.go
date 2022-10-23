package main

import "fmt"

func main() {
	// varで定義した変数は関数の外でも宣言可能
	var (
		i int = 1
		// デフォルトでは、float64が入る。float32で定義したい場合には、明示的にfloat32型をつけておく
	 f64 float64 = 1.2
	 s string = "test"
	 t, f bool = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 省略形は関数内でしか宣言できない
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}