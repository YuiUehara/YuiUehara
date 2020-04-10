package main
// 3章 関数と型 32p 【TRY】 スライスの利用 
// 正答例リンク"https://play.golang.org/p/FGuwWuFjQfO"

func main() {
	// タイトルから察するにスライスを使って
	// ↓こいつらを3つにまとめなければならない
	ns := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(ns); i++ { // lenで読み込んだnsの長さだけ、for文で繰り返させる。
// for文使うことがパッと出てこなかったので覚えておく。
/*
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12
*/
//		sum == ns[i]
		sum += ns[i] // sumに代入演算 ←要復習？
	}
//	sum := n1 + n2 + n3 + n4

	println(sum)
		}
