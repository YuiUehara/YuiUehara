package main

import "fmt"

func main() {
// 2章 基本構文 41p 【TRY】 奇数と偶数 
	for i := 1; i <＝ 100; i++ { /* 100まで繰り返せ */
			if i%2 == 0 {/* もしもiが偶数（2で割った数が0）だったら */
/*				fmt.Println(i + "-偶数")*/
				fmt.Printf("%d-偶数\n", i)
				/* %d は整数を10進法で出力するの意の指定子 */
			} else {
/*				fmt.Println(i + "-奇数")*/
				fmt.Printf("%d-奇数\n", i）
			}
		}

	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)

	println(
		"Hello,world"
	)
}
