package main

import "fmt"

// 一番最初に呼ばれる
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("Hello World")
}

// var i int = 1 // integer
// var f64 float64 = 1.5 // float64
// var s string = "test" // string
// var t bool = true // bool

// var (
//  i int = 1
//  f64 float64 = 1.5
// ) // 変数宣言まとめることが出来る

// func main {
//   // 型宣言を無しで記述(関数内でしか出来ない)
//   xi := 1
//   xf64 := 1.5
//   xt, xf := true, false
//   fmt.PrintIn("hello! golang") // fmtパッケージのPrintIn関数
//   fmt.Printf("%T/n", xf64) // 型を確認することが出来る関数
//   fmt.Printf("%T/n", xf64) // /nは改行
// }