package main
// 3章 関数と型 ５８p 【TRY】 ポインタ

func swap2(n, m *int){
	*n, *m = *m, *n // swap2が来たら、入れ替えるよ
}

func main() {
  n, m := 10, 20 // 数値は１０と20だよ
  swap2(&n, &m) // いま数字を入れ替えたよ
  println(n, m) // さぁ数字を確認してみたまえ
}

// 参考: 3章 関数と型 54p ポインタ
