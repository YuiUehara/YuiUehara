package main

// 4章 パッケージ 28p 【TRY】 ライブラリを取得してみよう

import (
	"fmt"

	//"github.com/tenntenn/greeting/v2"
	"greeting.go" //←これが無効になっていた
)

func main() {
	//	d := "こんにちは"()
	d := greeting.Do()
	fmt.Println(d)
}

// goのインポートでハマった
