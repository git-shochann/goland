package main

import "fmt"

// アドレスを引数に取る関数
func counter(adctr *int) {
	*adctr++ // アドレスに保存されている内容を読み取って,追加する
}

func freecounter(adctr *int, add int) {
	*adctr += add
}

func main() {
	ctr := 0
	fctr := 10

	for i:=0; i < 5; i++ {
		counter(&ctr) // 変数ctrのアドレスを渡す
		fmt.Printf("カウンタの値は%d\n", ctr)

		freecounter(&fctr, ctr)
	} 
}