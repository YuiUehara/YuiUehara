package main
// 3章 関数と型 65p 【TRY】 レシーバに変更を加える

type MyInt int
// Incメソッドは自身を1ずつ加算する

func (n *MyInt) Int() { *n++ } //←++が1ずつ加算させる演算子なのでメソッド

func main() {
	var n MyInt // 
	println(n) // 
	(n).Inc() // "."でメソッド（１ずつ加算）を召喚するレシーバ
	println(n) // 
}


// n.Inc undefined (type *MyInt has no field or method Inc)
// *MyInt にフィールドかメソッドがないよ

// レシーバに変更を加えるのはポインタ型
// おそらくポインタ型のメソッドリストを理解できてないから解けていない
// 参考リンク ポインタ型のメソッドリスト 
// "https://qiita.com/tenntenn/items/49bf8b5cc69c2fcfb627"
