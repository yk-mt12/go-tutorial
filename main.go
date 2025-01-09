package main

import (
	"fmt"
	"net/http"
	"math/rand"
)

func handler(w http.ResponseWriter, r *http.Request) {
	omikuji := []string{"大吉", "中吉", "小吉", "吉", "凶"}
	result := omikuji[rand.Intn(len(omikuji))]

	fmt.Fprintln(w, result)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
