package main
// 3章 関数と型 57p 【TRY】 複数戻り値の利用 

func swap(n, m int)(n2, m2 int) {
	n2, m2 = m, n
	return
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

// 参考: 3章 関数と型 48p 関数の定義
