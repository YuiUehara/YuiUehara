package main
// 3章 関数と型 10p 【TRY】 組み込み型（数値） 
// 設問内例文リンク"https://play.golang.org/p/pI0uTk6aZeM"

func main() {
//	var sum int　// sum は１０進法の整数です。
	var sum float32　// これなら小数点があっても使える。
	sum = 5 + 6 + 3 // 合計 
	avg := sum / 3 // sum を３で割った平均値
	if avg > 4.5 { // もし平均値が4.5より大きかったら
		println("good")
	}
}
// ./prog.go:7:9: constant 4.5 truncated to integer
// 型付き定数の値は常に定数型の値で正確でなければならない
// 小数が入るとintが使えない。
