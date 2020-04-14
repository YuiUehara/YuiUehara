package main

import ("fmt"
"math/rand"
"time"
)

func main() {
// 2章 基本構文 42p 【TRY】 おみくじプログラムを作ろう 
// 正答例リンク"https://play.golang.org/p/pruHZLxX8jc"
	t := time.Now().UnixNano() //現在時刻を数値で取得
	rand.Seed(t) //乱数の種を設定？ ←要復習
	n := rand.Intn(6) // 0-5までの乱数を発生
// println(s) //←多分、下にprintlnあるから要らない。

		switch n {
		case 5: // iが5の場合 
			fmt.Println("大吉")
		case 4,3: // iが4か3の場合  
			fmt.Println("中吉")
		case 2,1: // iが2か1の場合 
			fmt.Println("吉")
		default: // そうでない場合（数値は0~5の乱数なので0しかない）
			fmt.Println("凶")
		}
	}
/*
時間切れにつき未完成。似たようなの昔jsかPHPで作ったはずなので頑張ろう。多分岐なのでswhich文を使うのは合ってるはずだ。

完成。"}"が一つ多かったのが原因で随分と長く時間をかけてしまった。
*/
