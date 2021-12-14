package helloworld

import "fmt"

// 一番最初に呼ばれる
func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func hello() {
	bazz()
	fmt.Println("Hello World")
}