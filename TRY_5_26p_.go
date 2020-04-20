package main

// 5章 コマンドラインツール 26p 【TRY】 catコマンドを作ろう

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// p23より参照
	// パスを結合する
	path := filepath.Join("dir", "main.go")
	fmt.Println(path)
	// ファイル名を取得
	fmt.Println(filepath.Base(path))

	// p22 より参照
	// 標準入力から読み込む
	scanner := bufio.NewScanner(os.Stdin)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		//1行分を出力する
		fmt.Println(scanner.Text())
	}
	fmt.Println()
}
