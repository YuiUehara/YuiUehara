package main

func main() {
	var sum int　// sum は１０進法の整数です。
	sum = 5 + 6 + 3 // 合計 
	avg := sum / 3 // sum を３で割った平均値
	if avg > 4.5 { // もし平均値が4.5以上だったら
		println("good")
	}
}
