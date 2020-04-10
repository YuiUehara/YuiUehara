package main

func main() {
	for i := 1; 1 <= 100; i++{
		print(i)
		if(i%2 == 0 && i != 100) {
			println("-偶数")
		} else if(i%2 == 1) {
			println("-奇数")
		} else {
			println("-偶数")
			break;
		}
	}
}
	

// 何で１００まで繰り返さないのかわからないけど、とりあえず100でカウントが止まるようにした。
