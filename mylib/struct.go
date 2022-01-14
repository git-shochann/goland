package mylib

import "fmt"

type Mydata struct {
	// 変数 + 型
	Name string
	Data []int
}

func (md Mydata) PrintData() {
	fmt.Println("Name:", md.Name)
	fmt.Println("Data", md.Data)	
}

func init() {
	// 構造体の初期化
	taro := Mydata{
		"Hanako", []int{10,20,30,40,50},
	}
	// taroという変数は構造体Mydataから初期化している = 関数を所有しているので以下のように呼び出すことが可能
	taro.PrintData()
}