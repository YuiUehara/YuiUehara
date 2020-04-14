package main
// 3章 関数と型 65p 【TRY】 レシーバに変更を加える

type MyInt int
// Incメソッドは自身を1ずつ加算する

func (n *MyInt) Int() { *n++ } //←こいつが1ずつ加算

func main() {
	var n MyInt // 
	println(n) // 
	(n).Inc() // 
	println(n) // 
}


// n.Inc undefined (type *MyInt has no field or method Inc)
// *MyInt にフィールドかメソッドがないよ

// レシーバに変更を加えるのはポインタ型
