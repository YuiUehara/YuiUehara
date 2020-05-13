package main

// 10章 HTTPサーバとクライアント 21p 【TRY】 おみくじWebアプリの作成
// Webアプリ化してみよう
// HTTPサーバを作成する
// リクエストが来たらおみくじ の結果を返す
// 乱数の種は1回だけ初期化する
/// HTTPサーバを起動する前に初期化する

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	// 「io」はデータの読み書き、「util」はutillity(有用性)
	// 一つ一つの機能を組み合わせてエラーハンドリングとか実装できない(そもそも忘れちゃう)プログラマでも
	// 簡単にファイルの読み書きができるような機能をioutilというパッケージでを用意してあげたよ！便利でしょ！って感じ。
	// 入出力を抽象化したインターフェース、いわゆるファイル的な振る舞いをするものをまるっと同じように扱うための便利なもの

	"log" // 出力できるのはPrint, Fatal, Panicの３種類だけらしい
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

/* HTTPハンドラ */
// 引数に"レスポンスを書き込む先"と"クライアントからのリクエスト"を取る
// 第1引数はレスポンスを書き込む先
/// 書き込みにはfmtパッケージの関数などが使える
// 第2引数はクライアントからのリクエスト
//
/*
	func handler(w http.クライアントを書き込むWriter, r *http.クライアントからのリクエスト){
	fmt.Fprint(w, "レスポンスの書き込み")
*/

func main() {
	content, err := ioutil.ReadFile("n")

	t := time.Now().UnixNano() // 現在の時刻を数値で取得
	rand.Seed(t)               // 乱数の種を取得
	n := rand.Intn(9)          // 0-8までの乱数を発生させる
	//	for i := 1; i <= 3; i++{
	/*
		result = string
		defer print("今日の運勢は" + result + "です")
		switch n {
		case 0:
			result := "大吉"
		case 1:
			result := "中吉"
		case 2, 3:
			result := "吉"
		case 4, 5:
			result := "小吉"
		case 6, 7:
			result := "末吉"
		case 8:
			result := "凶"
		default:
			result := "大凶"
		}
	*/

	print("今日の運勢は『")
	//	defer fmt.Printf("』です。")
	switch n {
	case 0:
		fmt.Printf("大吉")
	case 1:
		fmt.Printf("中吉")
	case 2, 3:
		fmt.Printf("吉")
	case 4, 5:
		fmt.Printf("小吉")
	case 6, 7:
		fmt.Printf("末吉")
	case 8:
		fmt.Printf("凶")
	default:
		fmt.Printf("大凶")
	}

	println("』です。")

	if err != nil { //エラーが無以外だった場合
		log.Fatal(err)
	}
	fmt.Println(string(content))

	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		// 実行した時にエラーが無ければ↑が新規作成される。
		log.Fatal(err)
	}
	//
	fmt.Println(string(content))
	url := "http://example.com" // URLを定義 ここに自分で作ったおみくじのURL貼れば良いのでは？
	resp, _ := http.Get(url)    //
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	//
}
