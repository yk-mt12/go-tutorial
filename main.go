package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	 // リクエストパラメータを取得
	 name := r.FormValue("name")
	 if name == "" {
			 name = "ゲスト"
	 }

	 // 運勢をランダムに選択
	 omikuji := []string{"大吉", "中吉", "小吉", "凶"}
	 result := omikuji[0] // ここでは固定で "大吉" を返すようにしています

	 // レスポンスを生成
	 response := fmt.Sprintf("%sさんの運勢は「%s」です！", name, result)
	 fmt.Fprintln(w, response)
}

func main() {

	http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
