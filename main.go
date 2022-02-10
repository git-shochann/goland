package main

// 関数
// 無名関数
// 関数を返す関数

// 関数を引数に取る関数
// func CallFunction(f func()) {
// 	f() // ここで受け取って実行する
// }

// func main() {
// 	CallFunction(func (){
// 		fmt.Println("testing!")
// 	})
// }

// クロージャー
// 関数を返す関数 引数にstring, 戻り値もstringを返す
func later() func(string) string {
	var store string
}