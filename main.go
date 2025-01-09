package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")

	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(5 * time.Second)
	}()

	time.Sleep(6 * time.Second)
}
