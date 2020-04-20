package main

import (
	"fmt"
	"time"

	//"github.com/zuu-development/zuu-training/mod-training/greeting.go"
	//"github.com/tenntenn/greeting"　// こっちが最初の課題
	"github.com/tenntenn/greeting/v2" // こっちが最後の課題
)

func main() {
	//fmt.Println(greeting.Do()) // こっちは強制的にこんにちは
	fmt.Println(greeting.Do(time.Now())) // こっちは時間を読み取って回答してくれる
}
