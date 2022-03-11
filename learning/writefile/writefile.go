package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var input string
	var content string // 24行目以降で使いたいので、変数のスコープによりここに設定する -> 最初はベタ書きで大丈夫
	for {
		fmt.Println("何かを入力してください！, 's'で保存します。")
		fmt.Scanln(&input)
		
		if input == "s" {
			break
		}
		content += input // どんどん変数contentに追加していく
	}

	wdata := []byte(content) // []byteに変換
	err := ioutil.WriteFile("datafile.txt", wdata, 0777)
	if err != nil {
		fmt.Println("書き込みエラーです！")
		os.Exit(1)
	} else {
		fmt.Println(content)
		fmt.Println("以上ファイルに保存しました！")
	}
}

// Goの変数のスコープは{}で設定された範囲内