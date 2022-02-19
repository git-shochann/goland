package main

import "fmt"

func main() {
	commands := []string{"O", "S", "O", "S", "E", "Q"}

	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "N":
			fmt.Printf("[%s]新規文書を作成します。", commands[i])
		case "O":
			fmt.Printf("[%s]開くファイルを指定してください。", commands[i])
		case "S":
			fmt.Printf("[%s]ファイルを保存しました。", commands[i])
		case "Q":
			fmt.Printf("[%s]終了します。よろしいでしょうか。", commands[i])
		default:
			fmt.Printf("[%s]無効なコマンドです。", commands[i])
		}
	}
}
