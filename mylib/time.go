package mylib

import (
	"fmt"
	"time"
)

func init() {
	time := time.Now() // Nowメソッドの戻り値はTime型(構造体) -> その構造体が所有しているメソッドを使うことができる
	fmt.Printf("%T, %v", time, time) // time.Time timeパッケージ内のTime型, 2022-01-30 18:33:22.144008 +0900 JST
	unixTime := time.Unix() // Unix時間に変換
	fmt.Println(unixTime) // 1643534990
}
