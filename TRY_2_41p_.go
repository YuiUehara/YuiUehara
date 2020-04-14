package main

import ("fmt"
)

func main() {
// 2章 基本構文 41p 【TRY】 奇数と偶数 
// 正答例リンク"https://play.golang.org/p/pruHZLxX8jc"
	for i := 1; i <= 100; i++ { /* 100まで繰り返せ */
			if i%2 == 0 {// もしもiが偶数（2で割った数が0）だったら 
//				fmt.Println(i + "-偶数")
				fmt.Printf("%d-偶数\n", i)
				// %d は整数を10進法で出力するの意の指定子
			} else {
//				fmt.Println(i + "-奇数")
				fmt.Printf("%d-奇数\n", i)
			}
		}
	for i := 1; i <= 100; i++ {
			switch {
			case i%2 == 0: // iが偶数（2で割った数が0）の場合は 
				fmt.Printf("%d-偶数\n", i)
				break;
			default: // そうでない場合は
				fmt.Printf("%d-奇数\n", i)
				break;
			}
		}
}
// if~else文とswitch case defaultの違いとは？
// 二分岐だとif文、多分岐だとswitchが適切らしい。

/*
強いて言えば、数学の知識が足りていないのか？ 
演算子はある程度暗記しておかないと、使うたびにリスト確認することになってちょっと手間だ。
演算子を組み合わせるパズルは正直ちょっと楽しいけど、これにあまり時間をかけすぎるのはお仕事としてよろしくない。
エラーコード見てもわからなかったら、とりあえず全角を疑おう。
*/
