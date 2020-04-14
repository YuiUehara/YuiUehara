package main
// 数字を受け取って偶数か奇数を返す関数

func main() {
	for i := 1; i <= 100; i++{
		print(i)
		if(i%2 == 1) {
			println("-奇数")
		} else {
			println("-偶数")
		}
	}
}
	

/*
func main() {
	for i := 1; 1 <= 100; i++{
		print(i)
		if(i%2 == 0 && i != 100) { // ここを関数にして移動するする
//			println("-偶数")
		} else if(i%2 == 1) {
			println("-奇数")
		} else {
			println("-偶数")
		}
	}
}
*/	

// 何で１００まで繰り返さないのかわからないけど、力づくで100でカウントが止まるようにした。
