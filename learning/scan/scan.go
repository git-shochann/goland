package main

import "fmt"

func main() {
	var input string // メモリの確保 + 値はゼロ値で ""
	for {
		fmt.Print("何か入力してください！'q'で終了します")
		fmt.Scanln(&input) // アドレス？ 0x14000010230

		if input == "q" {
			break
		}
		fmt.Println(input, "が入力されました！")
	}
	fmt.Println("お疲れ様でした！")
}