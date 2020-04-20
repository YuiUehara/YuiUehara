package main

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
