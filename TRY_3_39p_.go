package main
// 3章 関数と型 39p 【TRY】 ユーザ定義型の利用 
// 正答例リンク"https://play.golang.org/p/FGuwWuFjQfO"
type Score struct { // "Score"の構造を設定します。
	UserID string // "UserID"は文字列です。
	GameID int // "GameID"は整数です。
	Point int
}

func main() {
}
		}
// "Score"の構造の中に何の情報を入れるかの設定、宣言？
// 思いつかなくて正当を確認したけどこれで本当に良いのか？？？

/*
 要は構造体"Score"を構成する要素で"ユーザーID"には文字列、"ゲームID"には整数、"ポイント"には整数が入るように決定します！って宣言させろという設問。
ここで宣言したを活かして、ステータス的な土台を設定していく設定。
もしかしたらこの後、会員サービス関係のプログラム作るときにこれのもっと複雑なやつを使うのかな？
*/
