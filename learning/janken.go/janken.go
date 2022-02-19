package main

import "fmt"

func main() {
	pattern := [3]string{"グー", "チョキ", "パー"}
	fmt.Println("***じゃんけん勝敗リスト***")

	for me := 0; me < 3; me++ { // 3回繰り返す
		fmt.Printf("わたしが%sのとき、\n", pattern[me])

		for you := 0; you < 3; you++ {
			score := (3 + me - you) % 3

			if score == 2 {
				fmt.Println("私の価値")
			} else if score == 1 {
				fmt.Println("私の負け")
			} else {
				fmt.Println("あいこ")
			}
		}
		fmt.Println()
	}
}
