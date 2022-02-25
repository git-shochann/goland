package main

import "fmt"

func counter() func() {
	
	ctr := 0
	fmt.Println("カウンタを初期化しました")
	fmt.Printf("カウンタのアドレスは%p\n", &ctr)
	
	// counter()が呼ばれてもまだ実行されない
	return func() {
		ctr++
		fmt.Printf("カウンタの値は%d、", ctr)
		fmt.Printf("カウンタのアドレスは%p\n", &ctr)
	}
}
func main() {
	countfunc := counter() // ここで実行して戻り値が関数
	countfunc()
	countfunc() 
}